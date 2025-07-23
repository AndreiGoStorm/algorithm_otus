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

var (
	ErrEqualZero = errors.New("count equal zero")
)

type FileImproveRLE struct {
	from   string
	to     string
	buffer []byte
	count  int
}

func NewFileImproveRLE(fileName string) *FileImproveRLE {
	return &FileImproveRLE{from: fileName}
}

func (f *FileImproveRLE) getToFileName(fromFile *os.File, ext string) string {
	fromFileExt := filepath.Ext(fromFile.Name())
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(fromFile.Name(), fromFileExt), ext)
}

func (f *FileImproveRLE) Compress() error {
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

func (f *FileImproveRLE) compressFile(from io.Reader, to io.Writer) error {
	buf := make([]byte, 1)

	var prev byte
	isFirst := true
	for {
		n, err := from.Read(buf)
		if n > 0 {
			b := buf[0]
			if isFirst {
				prev = b
				f.count = 1
				isFirst = false
				continue
			}

			if b == prev {
				if err := f.writeSingleBytes(to); err != nil {
					return err
				}
				if f.count < 127 {
					f.count++
				} else {
					if err := f.writeByte(to, prev); err != nil {
						return err
					}
				}
			} else {
				if f.count > 1 {
					if err := f.writeByte(to, prev); err != nil {
						return err
					}
				} else {
					f.buffer = append(f.buffer, prev)
					if len(f.buffer) == 127 {
						if err := f.writeSingleBytes(to); err != nil {
							return err
						}
					}
				}
				prev = b
				f.count = 1
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	if f.count > 1 {
		if err := f.writeByte(to, prev); err != nil {
			return err
		}
	} else {
		f.buffer = append(f.buffer, prev)
	}

	return f.writeSingleBytes(to)
}

func (f *FileImproveRLE) writeSingleBytes(to io.Writer) error {
	if len(f.buffer) > 0 {
		if _, err := to.Write([]byte{byte(int8(-len(f.buffer)))}); err != nil {
			return err
		}
		if _, err := to.Write(f.buffer); err != nil {
			return err
		}
		f.buffer = nil
	}
	return nil
}

func (f *FileImproveRLE) writeByte(to io.Writer, char byte) error {
	if _, err := to.Write([]byte{byte(f.count), char}); err != nil {
		return err
	}
	f.count = 1
	return nil
}

func (f *FileImproveRLE) Decompress() error {
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

func (f *FileImproveRLE) decompressFile(r io.Reader, w io.Writer) error {
	buf := make([]byte, 1)
	for {
		_, err := io.ReadFull(r, buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		count := int(int8(buf[0]))
		if count == 0 {
			return ErrEqualZero
		}

		if count > 0 {
			_, err = io.ReadFull(r, buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return err
			}

			block := bytes.Repeat([]byte{buf[0]}, count)
			if _, err := w.Write(block); err != nil {
				return err
			}
		} else {
			length := -count
			literalBytes := make([]byte, length)
			_, err = io.ReadFull(r, literalBytes)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return err
			}

			if _, err := w.Write(literalBytes); err != nil {
				return err
			}
		}
	}
	return nil
}
