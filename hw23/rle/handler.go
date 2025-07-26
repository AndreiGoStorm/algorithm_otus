package rle

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

type Handler struct {
	rf *os.File
	wf *os.File

	count int
	buf   []byte
}

func NewHandler(from string) (*Handler, error) {
	file, err := os.Open(from)
	if err != nil {
		return nil, err
	}
	return &Handler{rf: file}, nil
}

func (h *Handler) createWFile(to string) error {
	file, err := os.Create(to)
	if err != nil {
		return err
	}
	h.wf = file
	return nil
}

func (h *Handler) compress() (err error) {
	buf := make([]byte, 1)
	var prev byte
	isFirst := true
	for {
		if _, err = io.ReadFull(h.rf, buf); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		cur := buf[0]
		if isFirst {
			prev = cur
			h.count = 1
			isFirst = false
			continue
		}

		if prev, err = h.compressing(cur, prev); err != nil {
			return err
		}
	}

	if h.count > 1 {
		if err = h.writeByte(prev); err != nil {
			return err
		}
	} else {
		h.buf = append(h.buf, prev)
	}

	return h.writeSingleBytes()
}

func (h *Handler) compressing(cur, prev byte) (b byte, err error) {
	b = prev
	if cur == prev {
		if err = h.writeSingleBytes(); err != nil {
			return 0, err
		}
		if h.count < 127 {
			h.count++
		} else {
			if err = h.writeByte(prev); err != nil {
				return 0, err
			}
		}
	} else {
		if h.count > 1 {
			if err = h.writeByte(prev); err != nil {
				return 0, err
			}
		} else {
			h.buf = append(h.buf, prev)
			if len(h.buf) == 127 {
				if err = h.writeSingleBytes(); err != nil {
					return 0, err
				}
			}
		}
		b = cur
		h.count = 1
	}
	return b, nil
}

func (h *Handler) writeSingleBytes() error {
	if len(h.buf) > 0 {
		if _, err := h.wf.Write([]byte{byte(int8(-len(h.buf)))}); err != nil {
			return err
		}
		if _, err := h.wf.Write(h.buf); err != nil {
			return err
		}
		h.buf = nil
	}
	return nil
}

func (h *Handler) writeByte(char byte) error {
	if _, err := h.wf.Write([]byte{byte(h.count), char}); err != nil {
		return err
	}
	h.count = 1
	return nil
}

func (h *Handler) decompress() (err error) {
	buf := make([]byte, 1)
	for {
		if _, err = io.ReadFull(h.rf, buf); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		count := int(int8(buf[0]))
		if count == 0 {
			return errors.New("count equal zero")
		}

		done, err := h.decompressing(count, buf)
		if err != nil {
			return err
		}
		if done {
			break
		}
	}
	return nil
}

func (h *Handler) decompressing(count int, buf []byte) (done bool, err error) {
	if count > 0 {
		if _, err = io.ReadFull(h.rf, buf); err != nil {
			if errors.Is(err, io.EOF) {
				return true, nil
			}
			return false, err
		}

		block := bytes.Repeat([]byte{buf[0]}, count)
		if _, err = h.wf.Write(block); err != nil {
			return false, err
		}
	} else {
		length := -count
		literalBytes := make([]byte, length)

		if _, err = io.ReadFull(h.rf, literalBytes); err != nil {
			if errors.Is(err, io.EOF) {
				return true, nil
			}
			return false, err
		}

		if _, err = h.wf.Write(literalBytes); err != nil {
			return false, err
		}
	}
	return false, nil
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
