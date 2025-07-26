package main

type Compressor interface {
	Compress() error
	Decompress() error
	GetExtension() string
	GetDecompressedExtension() string
	Stat(bool)
}
