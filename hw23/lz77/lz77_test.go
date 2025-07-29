package lz77

import (
	"fmt"
	"os"
	"testing"

	"algorithm_otus/hw23/helpers"
	"github.com/stretchr/testify/require"
)

func TestLZ77(t *testing.T) {
	testDataFiles := (&helpers.Tester{}).GetTestDataFiles()
	for _, test := range testDataFiles {
		t.Run(fmt.Sprintf("file :%s", test.File), func(t *testing.T) {
			// Arrange
			compress, err := NewLZ77(test.File)
			require.NoError(t, err)
			compress.to = compress.fi.ReplaceExt(compress.from, compress.GetExtension())
			err = compress.h.createWFile(compress.to)
			require.NoError(t, err)

			decompress, err := NewLZ77(compress.to)
			require.NoError(t, err)

			// Act
			err = compress.Compress()
			require.NoError(t, err)
			err = decompress.Decompress()
			require.NoError(t, err)

			// Assert
			content, err := os.ReadFile(test.File)
			require.NoError(t, err)
			_, err = os.ReadFile(compress.to)
			require.NoError(t, err)

			decompressedContent, err := os.ReadFile(decompress.to)
			require.NoError(t, err)
			require.Equal(t, string(content), string(decompressedContent))

			//require.NoError(t, os.Remove(compress.to))
			//require.NoError(t, os.Remove(decompress.to))
		})
	}
}
