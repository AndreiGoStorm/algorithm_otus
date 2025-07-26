package huffman

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Writer struct {
	file *os.File
	bw   *BitWriter
}

func NewWriter(from, ext string) (w *Writer, err error) {
	w = &Writer{}
	w.file, err = os.Create(w.getWriteFileName(from, ext))
	if err != nil {
		return nil, err
	}
	w.bw = NewBitWriter(w.file)
	return w, nil
}

func (w *Writer) getWriteFileName(name, ext string) string {
	fromFileExt := filepath.Ext(name)
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(name, fromFileExt), ext)
}

func (w *Writer) writeFrequency(freq map[byte]int, total uint32) error {
	if err := binary.Write(w.file, binary.LittleEndian, uint16(len(freq))); err != nil {
		return err
	}
	for b, f := range freq {
		if err := binary.Write(w.file, binary.LittleEndian, b); err != nil {
			return err
		}
		if err := binary.Write(w.file, binary.LittleEndian, uint32(f)); err != nil {
			return err
		}
	}
	return binary.Write(w.file, binary.LittleEndian, total)
}

func (w *Writer) writeCodeTable(codeTable map[byte]string, r *Reader) (err error) {
	defer w.close()
	return r.encodeBytesAndWrite(codeTable, w.bw)
}

func (w *Writer) writeBytes(bytes []byte) error {
	buf := bufio.NewWriter(w.file)
	_, err := buf.Write(bytes)
	if err != nil {
		return err
	}
	return buf.Flush()
}

func (w *Writer) close() {
	if w.file != nil {
		err := w.file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}
