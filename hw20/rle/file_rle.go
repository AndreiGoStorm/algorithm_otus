package hw20rle

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileRLE struct {
	from string
	to   string
}

func NewFileRLE(fileName string) *FileRLE {
	return &FileRLE{from: fileName}
}

func (f *FileRLE) Compress() error {
	fromFile, err := os.Open(f.from)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	f.to = f.getToFileName(fromFile, "rle")

	toFile, err := os.Create(f.to)
	if err != nil {
		return err
	}
	defer toFile.Close()

	return f.compressFile(fromFile, toFile)
}

func (f *FileRLE) getToFileName(fromFile *os.File, ext string) string {
	fromFileExt := filepath.Ext(fromFile.Name())
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(fromFile.Name(), fromFileExt), ext)
}

func (f *FileRLE) compressFile(fromFile *os.File, toFile *os.File) error {
	buf := make([]byte, 1)

	var prev byte
	count := 0
	isFirst := true
	for {
		n, err := fromFile.Read(buf)
		if n > 0 {
			b := buf[0]
			if isFirst {
				prev = b
				count = 1
				isFirst = false
				continue
			}

			if b == prev && count < 255 {
				count++
			} else {
				if _, err := toFile.Write([]byte{prev, byte(count)}); err != nil {
					return err
				}
				prev = b
				count = 1
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	_, err := toFile.Write([]byte{prev, byte(count)})
	return err
}

func (f *FileRLE) Decompress() error {
	fromFile, err := os.Open(f.from)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	f.to = f.getToFileName(fromFile, "copy")

	toFile, err := os.Create(f.to)
	if err != nil {
		return err
	}
	defer toFile.Close()

	return f.decompressFile(fromFile, toFile)
}

func (f *FileRLE) decompressFile(r io.Reader, w io.Writer) error {
	buf := make([]byte, 2)
	var (
		value byte
		count int
	)
	for {
		_, err := io.ReadFull(r, buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		value, count = buf[0], int(buf[1])
		block := bytes.Repeat([]byte{value}, count)
		if _, err := w.Write(block); err != nil {
			return err
		}
	}
	return nil
}
