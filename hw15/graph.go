package hw15

import (
	"fmt"
	"strings"
)

const (
	NONE    int = iota // вершина не обнаружена
	FOUND              // вершина обнаружена но не посещена
	VISITED            // вершина посещенная или обработанная
)

type Graph struct {
	matrix [][]int
	states []int
	path   *List
	stack  *Stack
}

func NewGraph(matrix [][]int) *Graph {
	states := make([]int, len(matrix))
	return &Graph{matrix: matrix, states: states, path: &List{}, stack: &Stack{}}
}

func (g *Graph) Kan() bool {
	var canSort bool
	var deleted int
	for {
		canSort = false
		deleted = 0
		for i := 0; i < len(g.matrix); i++ {
			if g.states[i] != NONE {
				deleted++
				continue
			}
			if g.calculateNotVisitedColumns(i) == 0 {
				g.states[i] = VISITED
				g.path.PushBack(i + 1)
				canSort = true
				break
			}
		}
		if !canSort {
			break
		}
	}
	return deleted == len(g.matrix)
}

func (g *Graph) calculateNotVisitedColumns(i int) int {
	sum := 0
	for j := 0; j < len(g.matrix); j++ {
		if g.states[j] != VISITED {
			sum += g.matrix[j][i]
		}
	}
	return sum
}

func (g *Graph) Demukron() bool {
	weight := make([]int, len(g.matrix))
	sum := g.calculateWeights(weight)

	var canSort bool
	for {
		canSort = false
		for i := 0; i < len(g.matrix); i++ {
			if weight[i] > 0 {
				continue
			}
			if g.states[i] != VISITED {
				g.states[i] = VISITED
				g.path.PushBack(i + 1)
				for j := 0; j < len(g.matrix); j++ {
					weight[j] -= g.matrix[i][j]
					sum -= g.matrix[i][j]
				}
				canSort = true
				break
			}
		}
		if !canSort {
			return false
		}
		if sum <= 0 {
			break
		}
	}

	for i := 0; i < len(g.matrix); i++ {
		if g.states[i] != VISITED {
			g.path.PushBack(i + 1)
		}
	}
	return true
}

func (g *Graph) calculateWeights(weight []int) int {
	sum := 0
	for i := 0; i < len(g.matrix); i++ {
		for j := 0; j < len(g.matrix); j++ {
			weight[i] += g.matrix[j][i]
			sum += g.matrix[j][i]
		}
	}
	return sum
}

func (g *Graph) Tarian() bool {
	for i := 0; i < len(g.matrix); i++ {
		if g.states[i] == NONE {
			if !g.TarianDFS(i) {
				return false
			}
		}
	}
	for !g.stack.IsEmpty() {
		g.path.PushBack(g.stack.Pop())
	}
	return true
}

func (g *Graph) TarianDFS(z int) bool {
	g.states[z] = FOUND
	for i := 0; i < len(g.matrix); i++ {
		if g.matrix[z][i] == 0 {
			continue
		}
		if g.states[i] == NONE {
			if !g.TarianDFS(i) {
				return false
			}
		}
	}
	g.states[z] = VISITED
	g.stack.Push(z + 1)
	return true
}

func (g *Graph) pathToString() string {
	var builder strings.Builder
	for cur := g.path.Head; cur != nil; cur = cur.Next {
		builder.WriteString(fmt.Sprintf("%d ", cur.Value.(int)))
	}
	return strings.Trim(builder.String(), " ")
}
