package hw20rle

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompressFileRLE(t *testing.T) {
	for _, test := range []struct {
		file     string
		take     int
		expected []byte
	}{
		{
			file:     "testdata_rle/compress0.txt",
			take:     0,
			expected: []byte{0x47, 0x1, 0x65, 0x1, 0x74, 0x2, 0x69, 0x1, 0x6e, 0x1, 0x67, 0x1, 0x20, 0x1, 0x53, 0x1, 0x74, 0x1, 0x61, 0x1, 0x72, 0x1, 0x74, 0x1, 0x21, 0x3},
		},
		{
			file:     "testdata_rle/compress1.txt",
			take:     0,
			expected: []byte{0xd0, 0x1, 0x9f, 0x1, 0xd1, 0x1, 0x80, 0x1, 0xd0, 0x1, 0xb8, 0x1, 0xd0, 0x1, 0xb2, 0x1, 0xd0, 0x1, 0xb5, 0x1, 0xd1, 0x1, 0x82, 0x1, 0x2e, 0x3},
		},
		{
			file:     "testdata_rle/compress2.txt",
			take:     32,
			expected: []byte{0xd0, 0x1, 0x9a, 0x1, 0xd0, 0x1, 0xbe, 0x1, 0xd0, 0x1, 0xb3, 0x1, 0xd0, 0x1, 0xb4, 0x1, 0xd0, 0x1, 0xb0, 0x1, 0x20, 0x1, 0xd0, 0x1, 0xb7, 0x1, 0xd0, 0x1, 0xb8, 0x1, 0xd0, 0x1},
		},
	} {
		t.Run(fmt.Sprintf("file :%s", test.file), func(t *testing.T) {
			// Arrange
			rle := NewFileRLE(test.file)

			// Act
			err := rle.Compress()

			// Assert
			require.NoError(t, err)
			var fromFileContent []byte
			if test.take == 0 {
				fromFileContent, err = os.ReadFile(rle.to)
			} else {
				fromFileContent, err = readBytes(rle.to, test.take)
			}
			require.NoError(t, err)
			require.Equal(t, test.expected, fromFileContent)
			require.NoError(t, os.Remove(rle.to))
		})
	}
}

func readBytes(path string, count int) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, count)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buf[:n], nil
}

func TestDecompressFileRLE(t *testing.T) {
	for _, test := range []struct {
		file     string
		expected string
	}{
		{
			file:     "testdata_rle/decompress0.rle",
			expected: "testdata_rle/compress0.txt",
		},
		{
			file:     "testdata_rle/decompress1.rle",
			expected: "testdata_rle/compress1.txt",
		},
		{
			file:     "testdata_rle/decompress2.rle",
			expected: "testdata_rle/compress2.txt",
		},
	} {
		t.Run(fmt.Sprintf("file :%s", test.file), func(t *testing.T) {
			// Arrange
			rle := NewFileRLE(test.file)

			// Act
			err := rle.Decompress()

			// Assert
			require.NoError(t, err)
			expectedFileContent, err := os.ReadFile(test.expected)
			require.NoError(t, err)
			decompressedFileContent, err := os.ReadFile(rle.to)
			require.NoError(t, err)
			require.Equal(t, expectedFileContent, decompressedFileContent)
			require.NoError(t, os.Remove(rle.to))
		})
	}
}
