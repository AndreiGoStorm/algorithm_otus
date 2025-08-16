package hw17

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlgorithms(t *testing.T) {
	t.Run("Floid Warshell Algorithm", func(t *testing.T) {
		// Arrange
		graph := createFloidWarshellGraph()

		// Act
		ways, next := graph.FloidWarshell()

		// Assert
		graph.PrintFW(ways, next)
		require.Equal(t, 5, len(ways))
	})

	t.Run("Deikstra Algorithm", func(t *testing.T) {
		// Arrange
		graph := createDeikstraGraph()

		// Act
		for vertex := 1; vertex <= graph.size; vertex++ {
			ways := graph.Deikstra(vertex)
			_ = ways
			fmt.Printf("Минимальные пути из вершины %d:\n", vertex)
			for i := 1; i <= graph.size; i++ {
				fmt.Printf("%d ", ways[i])
			}
			fmt.Println()
		}

		// Assert
		require.Equal(t, 6, graph.size)
	})
}

func createFloidWarshellGraph() *Graph {
	matrix := map[int][]Vector{
		1: {{2, -2}, {3, 5}, {4, 7}},
		2: {{3, 6}, {4, 8}},
		3: {{1, -1}},
		4: {{2, 3}, {3, -4}},
	}

	return NewGraph(matrix)
}

func createDeikstraGraph() *Graph {
	matrix := map[int][]Vector{
		1: {{2, 7}, {3, 9}, {6, 14}},
		2: {{1, 7}, {3, 10}, {4, 15}},
		3: {{1, 9}, {2, 10}, {4, 11}, {6, 2}},
		4: {{2, 15}, {3, 11}, {5, 6}},
		5: {{4, 6}, {6, 9}},
		6: {{1, 14}, {3, 2}, {5, 9}},
	}

	return NewGraph(matrix)
}
