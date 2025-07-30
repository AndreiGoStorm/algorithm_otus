package hw16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTopologyAlgorithms(t *testing.T) {
	t.Run("Kan topology algorithm", func(t *testing.T) {
		// Arrange
		g := NewGraph(getGraphMatrix())

		// Act
		g.Kruskal()

		// Assert
		require.Equal(t, 12, g.minWeight)
	})
}

func getGraphMatrix() map[int][]Vector {
	return map[int][]Vector{
		1: {{2, 2}, {3, 2}, {5, 1}, {7, 3}},
		2: {{1, 2}, {3, 3}, {4, 3}},
		3: {{1, 2}, {2, 3}},
		4: {{2, 3}, {5, 2}},
		5: {{1, 1}, {4, 2}, {6, 4}},
		6: {{5, 4}, {7, 2}},
		7: {{1, 3}, {6, 2}},
	}
}
