package main

type Compressor interface {
	Compress() error
	Decompress() error
	GetName() string
	GetExtension() string
	GetDecompressExtension() string
	GetFileTo() string
}
