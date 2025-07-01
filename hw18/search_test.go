package hw18

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		text     string
		mask     string
		expected int
	}{
		{text: "BC-ABC-BC-AC-ABCBC-ABC-BC-C-ABC-ABC-BC-C-ABC", mask: "BC-ABC-BC-C-ABC", expected: 17},
		{text: "KOLKOKOLOKOLL", mask: "KOLOKOL", expected: 6},
		{text: "ABCDEFGHIJ", mask: "KOL", expected: -1},
		{text: "Новый способ", mask: "Новый", expected: 1},
		{text: "Русский текст", mask: "ск", expected: 4},
		{text: "кажется, что можно бы найти какой-то другой способ", mask: "другой", expected: 38},
		{text: "1349593113521355842359749613194152", mask: "1352135", expected: 9},
	}

	for _, tc := range tests {
		t.Run("full scan search", func(t *testing.T) {
			// Act
			pos, cmp := fullScanSearch(tc.text, tc.mask)

			// Assert
			require.Equal(t, pos, tc.expected)
			fmt.Printf("FullScanSearch cmp: %d - %s\n", cmp, tc.mask)
		})
	}

	for _, tc := range tests {
		t.Run("back scan search", func(t *testing.T) {
			// Act
			pos, cmp := backScanSearch(tc.text, tc.mask)

			// Assert
			require.Equal(t, pos, tc.expected)
			fmt.Printf("BackScanSearch cmp: %d - %s\n", cmp, tc.mask)
		})
	}

	for _, tc := range tests {
		t.Run("boyer moore horspool search shift 1", func(t *testing.T) {
			// Act
			pos, cmp := boyerMooreHorspoolShift1(tc.text, tc.mask)

			// Assert
			require.Equal(t, pos, tc.expected)
			fmt.Printf("BoyerMooreHorspool1 cmp: %d - %s\n", cmp, tc.mask)
		})
	}

	for _, tc := range tests {
		t.Run("boyer moore horspool search", func(t *testing.T) {
			// Act
			pos, cmp := boyerMooreHorspool(tc.text, tc.mask)

			// Assert
			require.Equal(t, pos, tc.expected)
			fmt.Printf("BoyerMooreHorspool cmp: %d - %s\n", cmp, tc.mask)
		})
	}

	for _, tc := range tests {
		t.Run("boyer moore search", func(t *testing.T) {
			// Act
			pos, cmp := boyerMoore(tc.text, tc.mask)

			// Assert
			require.Equal(t, pos, tc.expected)
			fmt.Printf("BoyerMoore cmp: %d - %s\n", cmp, tc.mask)
		})
	}
}
