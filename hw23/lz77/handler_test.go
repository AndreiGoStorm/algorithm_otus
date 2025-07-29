package lz77

import (
	"algorithm_otus/hw23/helpers"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWords(t *testing.T) {
	tst := helpers.NewTester()
	for _, tc := range tst.GetTestTextData() {
		t.Run("build frequency table", func(t *testing.T) {
			// Arrange
			file, _ := os.CreateTemp(tst.TempPath, tst.TempExt)
			defer os.Remove(file.Name())
			file.WriteString(tc.Text)

			compress, err := NewLZ77(file.Name())
			require.NoError(t, err)

			// Act
			err = compress.Compress()
			require.NoError(t, err)

			// Assert
			decompress, err := NewLZ77(compress.to)
			require.NoError(t, err)
			err = decompress.Decompress()
			require.NoError(t, err)

			content, err := os.ReadFile(compress.from)
			require.NoError(t, err)
			_, err = os.ReadFile(compress.to)
			require.NoError(t, err)

			decompressedContent, err := os.ReadFile(decompress.to)
			require.NoError(t, err)
			require.Equal(t, string(content), string(decompressedContent))

			require.NoError(t, os.Remove(compress.to))
			require.NoError(t, os.Remove(decompress.to))
		})
	}
}
