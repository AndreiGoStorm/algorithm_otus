package huffman

import (
	"io"
)

type BitWriter struct {
	w      io.Writer
	buffer byte
	count  uint8
}

func NewBitWriter(w io.Writer) *BitWriter {
	return &BitWriter{w: w}
}

func (bw *BitWriter) writeBit(bit byte) error {
	bw.buffer = (bw.buffer << 1) | bit
	bw.count++

	if bw.count == 8 {
		_, err := bw.w.Write([]byte{bw.buffer})
		if err != nil {
			return err
		}
		bw.buffer = 0
		bw.count = 0
	}
	return nil
}

func (bw *BitWriter) write(value string) error {
	var bit byte
	for _, c := range value {
		if c == '1' {
			bit = 1
		} else {
			bit = 0
		}

		err := bw.writeBit(bit)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bw *BitWriter) flush() error {
	if bw.count == 0 {
		return nil
	}
	bw.buffer <<= 8 - bw.count

	_, err := bw.w.Write([]byte{bw.buffer})
	bw.buffer = 0
	bw.count = 0
	return err
}
