package lz77

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

type Handler struct {
	rf     *os.File
	wf     *os.File
	buffer []byte
	rb     *RingBuffer
}

func NewHandler(from string) (*Handler, error) {
	file, err := os.Open(from)
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, windowSize)
	rb := NewRingBuffer(windowSize)
	return &Handler{rf: file, buffer: buffer, rb: rb}, nil
}

func (h *Handler) createWFile(to string) error {
	file, err := os.Create(to)
	if err != nil {
		return err
	}
	h.wf = file
	return nil
}

const windowSize = 4096
const lookaheadSize = 18

func (h *Handler) compress() error {
	if err := h.writeSize(); err != nil {
		return err
	}
	carry := []byte{}

	for {
		n, err := h.rf.Read(h.buffer)
		if n == 0 && err == io.EOF {
			if len(carry) > 0 {
				pos := 0
				for pos < len(carry) {
					bestLength, bestOffset := h.findMatches(carry, pos)

					var next byte
					if pos+bestLength < len(carry) {
						next = carry[pos+bestLength]
					} else {
						next = 0
					}

					if err = h.Write(bestOffset, bestLength, next); err != nil {
						return err
					}
					pos = h.updateWindow(carry, pos, bestLength+1)
				}
			}
			break
		}
		if err != nil && err != io.EOF {
			return err
		}

		lookahead := h.mergeCarry(carry, n)
		pos := 0
		for pos < len(lookahead) {
			bestLength, bestOffset := h.findMatches(lookahead, pos)
			if bestLength < 2 {
				bestLength = 0
				bestOffset = 0
			}
			if pos+bestLength >= len(lookahead) {
				carry = lookahead[pos:]
				break
			}

			next := lookahead[pos+bestLength]

			//fmt.Printf("%d\t%d\t%c:%[3]d\n", bestOffset, bestLength, next)
			if err = h.Write(bestOffset, bestLength, next); err != nil {
				return err
			}
			pos = h.updateWindow(lookahead, pos, bestLength+1)
		}

		if pos >= len(lookahead) {
			carry = nil
		}
	}

	return nil
}

func (h *Handler) writeSize() error {
	info, err := os.Stat(h.rf.Name())
	if err != nil {
		return err
	}
	if err = binary.Write(h.wf, binary.LittleEndian, uint64(info.Size())); err != nil {
		return err
	}
	return nil
}

func (h *Handler) mergeCarry(carry []byte, n int) (merged []byte) {
	if len(carry) > 0 {
		merged = make([]byte, len(carry)+n)
		copy(merged, carry)
		copy(merged[len(carry):], h.buffer[:n])
		return
	}
	merged = make([]byte, n)
	copy(merged, h.buffer[:n])
	return
}

func (h *Handler) findMatches(lookahead []byte, pos int) (bestLength, bestOffset int) {
	for i := 0; i < h.rb.len; i++ {
		length := 0
		for length < lookaheadSize &&
			pos+length < len(lookahead) &&
			i+length < h.rb.len &&
			lookahead[pos+length] == h.rb.Get(i+length) {
			length++
		}
		if length > 0 && length >= bestLength {
			bestLength = length
			bestOffset = h.rb.len - i
		}
	}
	return
}

func (h *Handler) updateWindow(lookahead []byte, pos, toAdd int) int {
	if toAdd > len(lookahead)-pos {
		toAdd = len(lookahead) - pos
	}
	h.rb.PushBytes(lookahead[pos : pos+toAdd])
	pos += toAdd
	return pos
}

func (h *Handler) Write(offset, length int, next byte) error {
	var buf [4]byte
	binary.LittleEndian.PutUint16(buf[0:2], uint16(offset))
	buf[2] = byte(length)
	buf[3] = next

	_, err := h.wf.Write(buf[:])
	return err
}

func (h *Handler) decompress() error {
	var fullSize uint64
	if err := binary.Read(h.rf, binary.LittleEndian, &fullSize); err != nil {
		return err
	}

	buf := make([]byte, 4)
	var written uint64 = 0

	for written < fullSize {
		_, err := io.ReadFull(h.rf, buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			if errors.Is(err, io.ErrUnexpectedEOF) {
				return fmt.Errorf("unexpected EOF: %w", err)
			}
			return err
		}

		offset := int(binary.LittleEndian.Uint16(buf[0:2]))
		length := int(buf[2])
		next := buf[3]

		if offset > h.rb.len {
			return fmt.Errorf("wrong offset: %d > window size %d", offset, h.rb.len)
		}

		if length > 0 {
			bytes, err := h.rb.GetBytesFromOffset(offset, length)
			if err != nil {
				return err
			}

			remain := fullSize - written
			if remain == 0 {
				break
			}
			if uint64(len(bytes)) > remain {
				bytes = bytes[:remain]
			}
			if _, err = h.wf.Write(bytes); err != nil {
				return err
			}
			h.rb.PushBytes(bytes)
			written += uint64(len(bytes))
		}

		if written < fullSize {
			if _, err := h.wf.Write([]byte{next}); err != nil {
				return err
			}
			h.rb.Push(next)
			written++
		}
	}

	return nil
}

func (h *Handler) close() {
	if h.rf != nil {
		if err := h.rf.Close(); err != nil {
			fmt.Println(err)
		}
	}
	if h.wf != nil {
		if err := h.wf.Close(); err != nil {
			fmt.Println(err)
		}
	}
}
