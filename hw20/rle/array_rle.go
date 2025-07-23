package hw20rle

import "errors"

type ArrayRLE struct {
	in  []byte
	out []byte
}

var (
	ErrEmptyInArray  = errors.New("empty in array")
	ErrWrongOutArray = errors.New("wrong out array")
)

func (a *ArrayRLE) Compress() error {
	if len(a.in) == 0 {
		return ErrEmptyInArray
	}

	a.out = make([]byte, 0, len(a.in))
	prev := a.in[0]
	count := byte(1)
	for i := 1; i < len(a.in); i++ {
		if a.in[i] == prev && count < 255 {
			count++
		} else {
			a.out = append(a.out, prev, count)
			prev = a.in[i]
			count = byte(1)
		}
	}

	a.out = append(a.out, prev, count)
	return nil
}

func (a *ArrayRLE) Decompress() error {
	if len(a.in) < 2 {
		return ErrWrongOutArray
	}

	a.out = make([]byte, 0, len(a.in))
	for i := 0; i < len(a.in); i += 2 {
		value := a.in[i]
		count := a.in[i+1]
		for count > 0 {
			a.out = append(a.out, value)
			count--
		}
	}
	return nil
}
