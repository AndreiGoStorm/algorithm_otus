package hw15

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTopologyAlgorithms(t *testing.T) {
	t.Run("Kan topology algorithm", func(t *testing.T) {
		// Arrange
		g := NewGraph(createGraphMatrix())

		// Act
		found := g.Kan()

		// Assert
		require.True(t, found)
		require.Equal(t, g.pathToString(), "1 3 2 4 6 5")
	})

	t.Run("Demukron topology algorithm", func(t *testing.T) {
		// Arrange
		g := NewGraph(createGraphMatrix())

		// Act
		found := g.Demukron()

		// Assert
		require.True(t, found)
		require.Equal(t, g.pathToString(), "1 3 2 4 6 5")
	})

	t.Run("Tarian topology algorithm", func(t *testing.T) {
		// Arrange
		g := NewGraph(createGraphMatrix())

		// Act
		found := g.Tarian()

		// Assert
		require.True(t, found)
		require.Equal(t, g.pathToString(), "3 1 4 2 6 5")
	})
}

func createGraphMatrix() [][]int {
	return [][]int{
		{0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
	}
}
