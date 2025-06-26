package hw13

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrefixTree(t *testing.T) {
	t.Run("insert item to prefix tree", func(t *testing.T) {
		// Arrange
		pt := NewPrefixTree()

		// Act
		pt.Insert("cat", "gray")

		// Assert
		value := pt.Search("cat")
		require.Equal(t, value, "gray")
	})

	t.Run("search not existed key", func(t *testing.T) {
		// Arrange
		pt := NewPrefixTree()
		pt.Insert("cat", "gray")

		// Act
		value := pt.Search("card")

		// Assert
		require.Nil(t, value)
	})

	t.Run("delete item in prefix tree", func(t *testing.T) {
		// Arrange
		pt := NewPrefixTree()

		// Act
		pt.Insert("dog", "black")
		pt.Insert("fox", "orange")
		pt.Insert("cat", "white")
		pt.Delete("cat")

		// Assert
		value := pt.Search("cat")
		require.Nil(t, value)
	})
}
