package huffman

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHuffman(t *testing.T) {
	for _, test := range []struct {
		file string
	}{
		{file: "../testdata/alice.txt"},
		{file: "../testdata/text_en.txt"},
		{file: "../testdata/text_pl.txt"},
		{file: "../testdata/text_ru.txt"},
		{file: "../testdata/archive.zip"},
		//{file: "../testdata/image.jpg"},
		//{file: "../testdata/till_down.mp3"},
		//{file: "../testdata/video.mp4"},
	} {
		t.Run(fmt.Sprintf("file :%s", test.file), func(t *testing.T) {
			// Arrange
			compress, err := NewHuffman(test.file)
			require.NoError(t, err)
			err = compress.createWriter(compress.GetExtension())
			require.NoError(t, err)
			compressFile := compress.w.file.Name()

			decompress, err := NewHuffman(compressFile)
			require.NoError(t, err)

			// Act
			err = compress.Compress()
			require.NoError(t, err)
			err = decompress.Decompress()
			require.NoError(t, err)

			// Assert
			content, err := os.ReadFile(test.file)
			require.NoError(t, err)
			_, err = os.ReadFile(compressFile)
			require.NoError(t, err)

			decompressFile := decompress.w.file.Name()
			decompressedContent, err := os.ReadFile(decompressFile)
			require.NoError(t, err)
			require.Equal(t, string(decompressedContent), string(content))

			require.NoError(t, os.Remove(compressFile))
			require.NoError(t, os.Remove(decompressFile))
		})
	}
}
