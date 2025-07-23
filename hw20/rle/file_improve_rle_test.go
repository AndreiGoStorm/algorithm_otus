package hw20rle

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	tempPath      = "/tmp/"
	tempExtension = "*.txt"
)

func TestCompressImproveRLEFileByString(t *testing.T) {
	for _, test := range []struct {
		str      string
		expected []byte
	}{
		{
			str:      "D",
			expected: []byte{0xff, 0x44},
		},
		{
			str:      "ADDBBBCD",
			expected: []byte{0xff, 0x41, 0x2, 0x44, 0x3, 0x42, 0xfe, 0x43, 0x44},
		},
		{
			str:      "AAABCEUU",
			expected: []byte{0x3, 0x41, 0xfd, 0x42, 0x43, 0x45, 0x2, 0x55},
		},
		{
			str:      "ABCDEFGHIJK",
			expected: []byte{0xf5, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b},
		},
		{
			str:      "AABCDEK",
			expected: []byte{0x2, 0x41, 0xfb, 0x42, 0x43, 0x44, 0x45, 0x4b},
		},
		{
			str:      "AAAAAAABBBC",
			expected: []byte{0x7, 0x41, 0x3, 0x42, 0xff, 0x43},
		},
	} {
		t.Run(fmt.Sprintf("file :%s", test.str), func(t *testing.T) {
			// Arrange
			file, _ := os.CreateTemp(tempPath, tempExtension)
			defer os.Remove(file.Name())
			file.WriteString(test.str)

			rle := NewFileImproveRLE(file.Name())

			// Act
			err := rle.Compress()

			// Assert
			require.NoError(t, err)
			compressedFileContent, err := os.ReadFile(rle.to)
			require.NoError(t, err)
			require.Equal(t, test.expected, compressedFileContent)
			require.NoError(t, os.Remove(rle.to))
		})
	}
}

func TestCompressFileImproveRLE(t *testing.T) {
	for _, test := range []struct {
		file     string
		expected []byte
	}{
		{
			file:     "testdata_improve_rle/compress0.txt",
			expected: []byte{0x7f, 0x41, 0xff, 0x41},
		},
		{
			file:     "testdata_improve_rle/compress1.txt",
			expected: []byte{0x7f, 0x41, 0xff, 0x41, 0x7f, 0x42, 0xff, 0x42},
		},
	} {
		t.Run(fmt.Sprintf("file :%s", test.file), func(t *testing.T) {
			// Arrange
			rle := NewFileImproveRLE(test.file)

			// Act
			err := rle.Compress()

			// Assert
			require.NoError(t, err)
			fromFileContent, err := os.ReadFile(rle.to)
			require.NoError(t, err)
			require.Equal(t, test.expected, fromFileContent)
			require.NoError(t, os.Remove(rle.to))
		})
	}
}

func TestDecompressFileImproveRLE(t *testing.T) {
	for _, test := range []struct {
		file     string
		expected string
	}{
		{
			file:     "testdata_improve_rle/decompress0.rle",
			expected: "testdata_improve_rle/compress0.txt",
		},
		{
			file:     "testdata_improve_rle/decompress1.rle",
			expected: "testdata_improve_rle/compress1.txt",
		},
		{
			file:     "testdata_improve_rle/decompress2.rle",
			expected: "testdata_improve_rle/compress2.txt",
		},
	} {
		t.Run(fmt.Sprintf("file :%s", test.file), func(t *testing.T) {
			// Arrange
			rle := NewFileImproveRLE(test.file)

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
