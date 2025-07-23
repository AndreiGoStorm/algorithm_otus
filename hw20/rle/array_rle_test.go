package hw20rle

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCompressArrayRLE(t *testing.T) {
	tests := []struct {
		name     string
		in       []byte
		expected []byte
	}{
		{
			name:     "one element",
			in:       []byte{5},
			expected: []byte{5, 1},
		},
		{
			name:     "data with no repeats",
			in:       []byte{1, 8, 3, 5},
			expected: []byte{1, 1, 8, 1, 3, 1, 5, 1},
		},
		{
			name:     "data with repeated elements",
			in:       []byte{5, 5, 5, 5, 5, 5, 6, 6, 6},
			expected: []byte{5, 6, 6, 3},
		},
		{
			name:     "data with mixed elements",
			in:       []byte{1, 1, 2, 2, 2, 4, 1, 7, 7, 7},
			expected: []byte{1, 2, 2, 3, 4, 1, 1, 1, 7, 3},
		},
	}

	// Arrange
	rle := &ArrayRLE{}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("compress for %s", tt.name), func(t *testing.T) {
			rle.in = tt.in

			// Act
			err := rle.Compress()

			// Assert
			require.NoError(t, err)
			require.Equal(t, tt.expected, rle.out)
		})
	}
}

func TestDecompressArrayRLE(t *testing.T) {
	tests := []struct {
		name     string
		in       []byte
		expected []byte
	}{
		{
			name:     "one element",
			in:       []byte{5, 1},
			expected: []byte{5},
		},
		{
			name:     "data with no repeats",
			in:       []byte{1, 1, 8, 1, 3, 1, 5, 1},
			expected: []byte{1, 8, 3, 5},
		},
		{
			name:     "data with repeated elements",
			in:       []byte{5, 6, 6, 3},
			expected: []byte{5, 5, 5, 5, 5, 5, 6, 6, 6},
		},
		{
			name:     "data with mixed elements",
			in:       []byte{1, 2, 2, 3, 4, 1, 1, 1, 7, 3},
			expected: []byte{1, 1, 2, 2, 2, 4, 1, 7, 7, 7},
		},
	}

	// Arrange
	rle := &ArrayRLE{}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("decompress for %s", tt.name), func(t *testing.T) {
			rle.in = tt.in

			// Act
			err := rle.Decompress()

			// Assert
			require.NoError(t, err)
			require.Equal(t, tt.expected, rle.out)
		})
	}
}
