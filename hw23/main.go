package main

import (
	"algorithm_otus/hw23/huffman"
	"flag"
	"fmt"
	"os"
)

var (
	from string
	lz77 bool
	rle  bool
	d    bool
)

func init() {
	flag.StringVar(&from, "from", "", "File to read from path. Required.")
	flag.BoolVar(&lz77, "lz77", false, "Mode to choose algorithm lz77. Huffman algorithm by default.")
	flag.BoolVar(&rle, "rle", false, "Mode to choose algorithm rle. Huffman algorithm by default.")
	flag.BoolVar(&d, "d", false, "Operation to decompress file. Compress operation by default.")
}

func main() {
	flag.Parse()

	if from == "" {
		err := fmt.Errorf("missing required flag: -from")
		fmt.Fprintln(os.Stderr, "Error:", err)
		flag.Usage()
		os.Exit(1)
	}

	c, err := getCompressor()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		flag.Usage()
		os.Exit(1)
	}

	if !d {
		err = c.Compress()
	} else {
		err = c.Decompress()
	}

	if err == nil {
		fmt.Println("________________________________________________")
		c.Stat(d)
	} else {
		fmt.Println(err)
	}
}

func getCompressor() (Compressor, error) {
	if lz77 {
		fmt.Println("LZ77 algorithm")
		//c, err := huffman.NewLZ(from)
		//return c, err
	}
	if rle {
		fmt.Println("RLE algorithm")
		//c, err := huffman.NewRLE(from)
		//return c, err
	}
	fmt.Println("Huffman algorithm")
	c, err := huffman.NewHuffman(from)
	return c, err
}
