package huffman

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var (
	tempPath      = "/tmp/"
	tempExtension = "*.txt"
)

func TestHuffmanTree(t *testing.T) {
	text := "ALICE IN WONDERLAND"
	tree := NewTree()

	t.Run("build frequency table", func(t *testing.T) {
		// Arrange
		file, _ := os.CreateTemp(tempPath, tempExtension)
		defer os.Remove(file.Name())
		file.WriteString(text)
		r, err := NewReader(file.Name())

		// Act
		require.NoError(t, err)
		frequency, err := r.prepareFrequency()

		// Assert
		expected := map[byte]int{
			'A': 2, //65
			'L': 2, //76
			'I': 2, //73
			'C': 1, //67
			'E': 2, //69
			' ': 2, //32
			'N': 3, //78
			'W': 1, //87
			'O': 1, //79
			'D': 2, //68
			'R': 1, //82
		}
		for key, value := range expected {
			require.Equal(t, value, frequency[key])
		}
		require.Equal(t, uint32(19), r.total)
	})

	t.Run("build huffman tree", func(t *testing.T) {
		// Arrange
		file, _ := os.CreateTemp(tempPath, tempExtension)
		defer os.Remove(file.Name())
		file.WriteString(text)
		r, err := NewReader(file.Name())

		// Act
		require.NoError(t, err)
		frequency, err := r.prepareFrequency()
		require.NoError(t, err)
		tree.BuildHuffmanTree(frequency)

		// Assert
		require.Equal(t, byte('A'), tree.root.right.right.right.right.value) // 1111
		require.Equal(t, byte('L'), tree.root.left.right.right.value)        // 011
		require.Equal(t, byte('I'), tree.root.left.right.left.value)         // 010
		require.Equal(t, byte('C'), tree.root.right.left.left.left.value)    // 1000
		require.Equal(t, byte('E'), tree.root.left.left.right.value)         // 001
		require.Equal(t, byte(' '), tree.root.right.right.right.left.value)  // 1110
		require.Equal(t, byte('N'), tree.root.right.right.left.value)        // 110
		require.Equal(t, byte('W'), tree.root.right.left.right.right.value)  // 1011
		require.Equal(t, byte('O'), tree.root.right.left.left.right.value)   // 1001
		require.Equal(t, byte('D'), tree.root.left.left.left.value)          // 000
		require.Equal(t, byte('R'), tree.root.right.left.right.left.value)   // 1010
		require.Equal(t, 19, tree.root.weight)
	})

	t.Run("build code table", func(t *testing.T) {
		// Arrange
		file, _ := os.CreateTemp(tempPath, tempExtension)
		defer os.Remove(file.Name())
		file.WriteString(text)
		r, err := NewReader(file.Name())

		// Act
		require.NoError(t, err)
		frequency, err := r.prepareFrequency()
		require.NoError(t, err)
		tree.BuildHuffmanTree(frequency)
		codeTable := tree.BuildCodeTable(tree.root)

		// Assert
		expected := map[byte]string{
			'A': "1111",
			'L': "011",
			'I': "010",
			'C': "1000",
			'E': "001",
			' ': "1110",
			'N': "110",
			'W': "1011",
			'O': "1001",
			'D': "000",
			'R': "1010",
		}
		require.Equal(t, len(expected), len(codeTable))
		for key, value := range expected {
			require.Equal(t, value, codeTable[key])
		}
	})
}
