package lz77

import (
	"algorithm_otus/hw23/helpers"
	"fmt"
)

type LZ77 struct {
	from string
	to   string

	h  *Handler
	fi *helpers.FileInfo
}

func NewLZ77(from string) (*LZ77, error) {
	h, err := NewHandler(from)
	if err != nil {
		return nil, err
	}
	return &LZ77{
		from: from, h: h, fi: helpers.NewFileInfo(),
	}, nil
}

func (lz *LZ77) Compress() error {
	defer lz.h.close()

	lz.to = lz.fi.ReplaceExt(lz.from, lz.GetExtension())
	if err := lz.h.createWFile(lz.to); err != nil {
		return err
	}

	return lz.h.compress()
}

func (lz *LZ77) Decompress() error {
	defer lz.h.close()

	lz.to = lz.fi.ReplaceExt(lz.from, lz.GetDecompressExtension())
	if err := lz.h.createWFile(lz.to); err != nil {
		return err
	}

	return lz.h.decompress()
}

func (lz *LZ77) GetName() string {
	return "LZ77 algorithm"
}

func (lz *LZ77) GetExtension() string {
	return "lz77"
}

func (lz *LZ77) GetDecompressExtension() string {
	return fmt.Sprintf("%s.dec", lz.GetExtension())
}

func (lz *LZ77) GetFileTo() string {
	return lz.to
}
