package huffman

import (
	"algorithm_otus/hw23/helpers"
	"fmt"
)

type Huffman struct {
	from string
	r    *Reader
	to   string
	w    *Writer

	t  *Tree
	fi *helpers.FileInfo
}

func NewHuffman(from string) (*Huffman, error) {
	h := &Huffman{}
	if err := h.createReader(from); err != nil {
		return nil, err
	}
	h.t = NewTree()
	h.fi = helpers.NewFileInfo()

	return h, nil
}

func (h *Huffman) createReader(from string) error {
	r, err := NewReader(from)
	if err != nil {
		return err
	}
	h.from = from
	h.r = r
	return nil
}

func (h *Huffman) createWriter(to string) error {
	w, err := NewWriter(to)
	if err != nil {
		return err
	}
	h.to = to
	h.w = w
	return nil
}

func (h *Huffman) Compress() error {
	defer h.close()

	if err := h.createWriter(h.fi.ReplaceExt(h.from, h.GetExtension())); err != nil {
		return err
	}

	frequency, err := h.r.prepareFrequency()
	if err != nil {
		return err
	}

	if err := h.w.writeFrequency(frequency, h.r.total); err != nil {
		return err
	}

	h.t.BuildHuffmanTree(frequency)
	codeTable := h.t.BuildCodeTable(h.t.root)

	return h.w.writeCodeTable(codeTable, h.r)
}

func (h *Huffman) Decompress() error {
	defer h.close()

	if err := h.createWriter(h.fi.ReplaceExt(h.from, h.GetDecompressExtension())); err != nil {
		return err
	}

	frequency, err := h.r.readFrequency()
	if err != nil {
		return err
	}
	h.t.BuildHuffmanTree(frequency)

	bytes, err := h.r.readCodeTable(h.t.root)
	if err != nil {
		return err
	}

	return h.w.writeBytes(bytes)
}

func (h *Huffman) close() {
	h.r.close()
	h.w.close()
}

func (h *Huffman) GetName() string {
	return "Huffman algorithm"
}

func (h *Huffman) GetExtension() string {
	return "huff"
}

func (h *Huffman) GetDecompressExtension() string {
	return fmt.Sprintf("%s.dec", h.GetExtension())
}

func (h *Huffman) GetFileTo() string {
	return h.to
}
