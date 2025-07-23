package hw20rle

type Compressor interface {
	Compress() error
	Decompress() error
}
