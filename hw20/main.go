package main

import (
	"flag"
	"fmt"
	"os"

	hw20rle "algorithm_otus/hw20/rle"
)

var (
	from string
	mode bool
	d    bool
)

func init() {
	flag.StringVar(&from, "from", "", "File to read from path. Required.")
	flag.BoolVar(&mode, "mode", false, "Mode to choose algorithm: simple rle or improved rle. Simple rle by default.")
	flag.BoolVar(&d, "d", false, "Operation to decompress file. Compress operation by default.")
}

func main() {
	flag.Parse()

	if from == "" {
		err := fmt.Errorf("missing required flag: -name")
		fmt.Fprintln(os.Stderr, "Error:", err)
		flag.Usage()
		os.Exit(1)
	}

	var rle hw20rle.Compressor
	if !mode {
		rle = hw20rle.NewFileRLE(from)
	} else {
		rle = hw20rle.NewFileImproveRLE(from)
	}

	var err error
	if !d {
		err = rle.Compress()
	} else {
		err = rle.Decompress()
	}

	if err != nil {
		fmt.Println(err)
	}
}
