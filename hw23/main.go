package main

import (
	"flag"
	"fmt"
	"os"

	"algorithm_otus/hw23/helpers"
	"algorithm_otus/hw23/huffman"
	"algorithm_otus/hw23/lz77"
	rl "algorithm_otus/hw23/rle"
)

var (
	from string
	lz   bool
	rle  bool
	d    bool
)

func init() {
	flag.StringVar(&from, "from", "", "File to read from path. Required.")
	flag.BoolVar(&lz, "lz", false, "Mode to choose algorithm lz77. Huffman algorithm by default.")
	flag.BoolVar(&rle, "rle", false, "Mode to choose algorithm rle. Huffman algorithm by default.")
	flag.BoolVar(&d, "d", false, "Action to decompress file. Compress action by default.")
}

func main() {
	flag.Parse()

	if from == "" {
		fail(fmt.Errorf("missing required flag: -from"))
	}

	c, err := getCompressor()
	if err != nil {
		fail(err)
	}

	if !d {
		err = c.Compress()
	} else {
		err = c.Decompress()
	}

	if err != nil {
		fail(err)
	}

	fmt.Println(c.GetName())
	fmt.Println("________________________________________________")

	stat := helpers.NewStat(d)
	if err = stat.Print(from, c.GetFileTo()); err != nil {
		fail(err)
	}
}

func getCompressor() (c Compressor, err error) {
	if lz {
		c, err = lz77.NewLZ77(from)
		return c, err
	}
	if rle {
		c, err = rl.NewRLE(from)
		return c, err
	}
	c, err = huffman.NewHuffman(from)
	return c, err
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	flag.Usage()
	os.Exit(1)
}
