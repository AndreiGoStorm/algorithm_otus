package huffman

import (
	"algorithm_otus/hw23/helpers"
	"fmt"
)

type Huffman struct {
	from string
	r    *Reader
	w    *Writer
	t    *Tree
}

func NewHuffman(from string) (*Huffman, error) {
	r, err := NewReader(from)
	if err != nil {
		return nil, err
	}

	return &Huffman{from: from, r: r, t: NewTree()}, nil
}

func (h *Huffman) createWriter(ext string) error {
	w, err := NewWriter(h.from, ext)
	if err != nil {
		return err
	}
	h.w = w
	return nil
}

func (h *Huffman) Compress() error {
	if err := h.createWriter(h.GetExtension()); err != nil {
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
	if err := h.createWriter(h.GetDecompressedExtension()); err != nil {
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

func (h *Huffman) GetExtension() string {
	return "huff"
}

func (h *Huffman) GetDecompressedExtension() string {
	return fmt.Sprintf("%s.dec", h.GetExtension())
}

func (h *Huffman) Stat(isDecompress bool) {
	fileHelper, err := helpers.NewFileInfo(h.r.file.Name())
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("%-15s %-20s| %6s\n", "Source file", fileHelper.Filename, fileHelper.HumanSize())

	err = fileHelper.SetPath(h.w.file.Name())

	message := "Compression"
	if isDecompress {
		message = "Decompression"
	}
	fmt.Printf("%-15s %-20s| %6s\n", message, fileHelper.Filename, fileHelper.HumanSize())
}
