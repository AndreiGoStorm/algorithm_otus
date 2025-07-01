package hw19

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected string
	}{
		{text: "ABABABCCBBABABCAB", pattern: "ABABC", expected: "2 10"},
		{text: "BABABABCCBBABABCAB", pattern: "ABA", expected: "1 3 11"},
		{text: "Searching Auto Algorithm", pattern: "ch", expected: "4"},
		{text: "Algorithm Knuth Morris Pratt", pattern: "th", expected: "6 13"},
	}

	for _, tc := range tests {
		t.Run("Search Auto Algorithm", func(t *testing.T) {
			// Act
			actual := Search(tc.text, tc.pattern)

			// Assert
			require.Equal(t, actual, tc.expected)
		})
	}

	for _, tc := range tests {
		t.Run("Algorithm Knuth Morris Pratt", func(t *testing.T) {
			// Act
			actual := KnuthMorrisPrattSearch(tc.text, tc.pattern)

			// Assert
			require.Equal(t, actual, tc.expected)
		})
	}
}
