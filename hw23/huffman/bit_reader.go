package huffman

import "io"

type BitReader struct {
	r      io.Reader
	buffer byte
	count  uint8
}

func NewBitReader(r io.Reader) *BitReader {
	return &BitReader{r: r}
}

func (br *BitReader) ReadBit() (byte, error) {
	if br.count == 0 {
		buf := make([]byte, 1)
		n, err := br.r.Read(buf)
		if err != nil {
			return 0, err
		}
		if n != 1 {
			return 0, io.ErrUnexpectedEOF
		}
		br.buffer = buf[0]
		br.count = 8
	}
	bit := (br.buffer >> 7) & 1
	br.buffer <<= 1
	br.count--
	return bit, nil
}
