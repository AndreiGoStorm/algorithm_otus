package lz77

import (
	"errors"
	"fmt"
)

type RingBuffer struct {
	buf    []byte
	cap    int
	len    int
	cursor int
}

func NewRingBuffer(cap int) *RingBuffer {
	return &RingBuffer{
		buf: make([]byte, cap),
		cap: cap,
	}
}

func (r *RingBuffer) Push(b byte) {
	r.buf[r.cursor] = b
	r.cursor = (r.cursor + 1) % r.cap
	if r.len < r.cap {
		r.len++
	}
}

func (r *RingBuffer) PushBytes(bytes []byte) {
	for i := 0; i < len(bytes); i++ {
		r.Push(bytes[i])
	}
}

func (r *RingBuffer) GetOldByte() (byte, error) {
	if r.len == 0 {
		return 0, fmt.Errorf("empty buffer")
	}
	start := (r.cursor + r.cap - r.len) % r.cap
	return r.buf[start], nil
}

func (r *RingBuffer) Get(i int) byte {
	if i < 0 || i >= r.len {
		panic("out of bounds")
	}
	start := (r.cursor + r.cap - r.len) % r.cap
	return r.buf[(start+i)%r.cap]
}

func (r *RingBuffer) IsEmpty() bool {
	return r.len == 0
}

func (r *RingBuffer) GetBytesFromOffset(offset, length int) ([]byte, error) {
	if offset <= 0 || offset > r.cap {
		return nil, errors.New("invalid offset")
	}
	if offset > r.len {
		return nil, errors.New("offset out of bounds")
	}

	start := (r.cursor - offset + r.cap) % r.cap
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = r.buf[(start+i)%r.cap]
	}
	return result, nil
}
