package hw17

import (
	"fmt"
	"math"
)

type Graph struct {
	matrix map[int][]Vector
	size   int
}

const INF = math.MaxInt

func NewGraph(matrix map[int][]Vector) *Graph {
	size := len(matrix)
	return &Graph{matrix: matrix, size: size}
}

func (g *Graph) FloidWarshell() ([][]int, [][]int) {
	ways, next := g.prepareFWWays()

	for k := 1; k <= g.size; k++ {
		for i := 1; i <= g.size; i++ {
			for j := 1; j <= g.size; j++ {
				if ways[i][k] == INF || ways[k][j] == INF {
					continue
				}
				if ways[i][k]+ways[k][j] < ways[i][j] {
					ways[i][j] = ways[i][k] + ways[k][j]
					next[i][j] = next[i][k]
				}
			}
		}
	}
	return ways, next
}

func (g *Graph) prepareFWWays() ([][]int, [][]int) {
	ways := make([][]int, g.size+1)
	next := make([][]int, g.size+1)
	for i := 1; i <= g.size; i++ {
		ways[i] = make([]int, g.size+1)
		next[i] = make([]int, g.size+1)
		for j := 1; j <= g.size; j++ {
			if i != j {
				ways[i][j] = INF
			}
			next[i][j] = -1
		}
	}

	for from, vectors := range g.matrix {
		for _, v := range vectors {
			ways[from][v.To] = v.Weight
			next[from][v.To] = v.To
		}
	}

	return ways, next
}

func (g *Graph) PrintFW(ways [][]int, next [][]int) {
	for i := 1; i <= g.size; i++ {
		for j := 1; j <= g.size; j++ {
			if i == j {
				continue
			}
			if ways[i][j] == INF {
				continue
			}
			fmt.Printf("%d to %d | weight = %2d | ", i, j, ways[i][j])
			g.printFWPath(next, i, j)
		}
	}
}

func (g *Graph) printFWPath(next [][]int, i, j int) {
	if next[i][j] == -1 {
		return
	}

	path := []int{i}
	for i != j {
		i = next[i][j]
		path = append(path, i)
	}

	for k := 0; k < len(path); k++ {
		if k > 0 {
			fmt.Print(" → ")
		}
		fmt.Print(path[k])
	}
	fmt.Println()
}

func (g *Graph) Deikstra(vertex int) []int {
	ways, visited := g.prepareDWays(vertex)

	for k := 1; k <= g.size; k++ {
		minimum := g.getMinimum(ways, visited)
		visited[minimum] = true

		for _, v := range g.matrix[minimum] {
			if visited[v.To] {
				continue
			}
			weight := ways[minimum] + v.Weight
			if weight < ways[v.To] {
				ways[v.To] = weight
			}
		}
	}
	return ways
}

func (g *Graph) prepareDWays(vertex int) ([]int, []bool) {
	ways := make([]int, g.size+1)
	visited := make([]bool, g.size+1)
	for i := 1; i <= g.size; i++ {
		ways[i] = INF
	}
	ways[vertex] = 0

	return ways, visited
}

func (g *Graph) getMinimum(ways []int, visited []bool) int {
	minimum := -1
	for i := 1; i <= g.size; i++ {
		if visited[i] {
			continue
		}
		if minimum == -1 || ways[i] < ways[minimum] {
			minimum = i
		}
	}
	return minimum
}
