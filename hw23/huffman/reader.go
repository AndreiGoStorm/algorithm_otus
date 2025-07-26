package huffman

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Reader struct {
	file  *os.File
	total uint32
	br    *BitReader
}

func NewReader(from string) (*Reader, error) {
	file, err := os.Open(from)
	if err != nil {
		return nil, err
	}
	return &Reader{file: file, br: NewBitReader(file)}, nil
}

func (r *Reader) prepareFrequency() (map[byte]int, error) {
	frequency := make(map[byte]int, 256)
	buf := make([]byte, 4096)

	for {
		n, err := r.file.Read(buf)
		if n > 0 {
			for _, b := range buf[:n] {
				frequency[b]++
				r.total++
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return frequency, nil
}

func (r *Reader) encodeBytesAndWrite(codeTable map[byte]string, bw *BitWriter) (err error) {
	if _, err = r.file.Seek(0, io.SeekStart); err != nil {
		return err
	}
	buf := make([]byte, 4096)
	for {
		n, err := r.file.Read(buf)
		if n > 0 {
			for _, b := range buf[:n] {
				code, ok := codeTable[b]
				if !ok {
					return fmt.Errorf("byte %v not found in code table", b)
				}
				if err := bw.write(code); err != nil {
					return err
				}
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return bw.flush()
}

func (r *Reader) readFrequency() (map[byte]int, error) {
	var symbolCount uint16
	if err := binary.Read(r.file, binary.LittleEndian, &symbolCount); err != nil {
		return nil, err
	}

	frequency := make(map[byte]int, symbolCount)
	var b byte
	var f uint32
	for i := 0; i < int(symbolCount); i++ {
		if err := binary.Read(r.file, binary.LittleEndian, &b); err != nil {
			return nil, err
		}
		if err := binary.Read(r.file, binary.LittleEndian, &f); err != nil {
			return nil, err
		}
		frequency[b] = int(f)
	}

	return frequency, nil
}

func (r *Reader) readCodeTable(root *Node) ([]byte, error) {
	var byteCount uint32
	if err := binary.Read(r.file, binary.LittleEndian, &byteCount); err != nil {
		return nil, err
	}

	current := root
	var output []byte
	for uint32(len(output)) < byteCount {
		bit, err := r.br.ReadBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if bit == 0 {
			current = current.left
		} else {
			current = current.right
		}

		if current.left == nil && current.right == nil {
			output = append(output, current.value)
			current = root
		}
	}
	return output, nil
}

func (r *Reader) close() {
	if r.file != nil {
		if err := r.file.Close(); err != nil {
			fmt.Println(err)
		}
	}
}
