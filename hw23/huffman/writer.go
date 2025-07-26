package huffman

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

type Writer struct {
	file *os.File
	bw   *BitWriter
}

func NewWriter(to string) (w *Writer, err error) {
	w = &Writer{}
	w.file, err = os.Create(to)
	if err != nil {
		return nil, err
	}
	w.bw = NewBitWriter(w.file)
	return w, nil
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
	return r.encodeBytesAndWrite(codeTable, w.bw)
}

func (w *Writer) writeBytes(bytes []byte) error {
	buf := bufio.NewWriter(w.file)
	if _, err := buf.Write(bytes); err != nil {
		return err
	}
	return buf.Flush()
}

func (w *Writer) close() {
	if w.file != nil {
		if err := w.file.Close(); err != nil {
			fmt.Println(err)
		}
	}
}
