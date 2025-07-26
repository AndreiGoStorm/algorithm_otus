package rle

import (
	"algorithm_otus/hw23/helpers"
	"fmt"
)

type RLE struct {
	from string
	to   string
	h    *Handler
	fi   *helpers.FileInfo
}

func NewRLE(from string) (*RLE, error) {
	h, err := NewHandler(from)
	if err != nil {
		return nil, err
	}
	return &RLE{
		from: from, h: h, fi: helpers.NewFileInfo(),
	}, nil
}

func (rle *RLE) Compress() error {
	defer rle.h.close()

	rle.to = rle.fi.ReplaceExt(rle.from, rle.GetExtension())
	if err := rle.h.createWFile(rle.to); err != nil {
		return err
	}

	return rle.h.compress()
}

func (rle *RLE) Decompress() error {
	defer rle.h.close()

	rle.to = rle.fi.ReplaceExt(rle.from, rle.GetDecompressExtension())
	if err := rle.h.createWFile(rle.to); err != nil {
		return err
	}

	return rle.h.decompress()
}

func (rle *RLE) GetName() string {
	return "RLE algorithm"
}

func (rle *RLE) GetExtension() string {
	return "rle"
}

func (rle *RLE) GetDecompressExtension() string {
	return fmt.Sprintf("%s.dec", rle.GetExtension())
}

func (rle *RLE) GetFileTo() string {
	return rle.to
}
