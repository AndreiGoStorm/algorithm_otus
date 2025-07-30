package hw16

import (
	"fmt"
	"sort"
)

type Graph struct {
	matrix map[int][]Vector
	size   int

	parents   []int
	edges     []Edge
	q         *Queue
	minWeight int
}

func NewGraph(matrix map[int][]Vector) *Graph {
	g := &Graph{matrix: matrix, size: len(matrix)}
	g.prepare()
	return g
}

func (g *Graph) prepare() {
	g.parents = make([]int, g.size+1)
	g.edges = make([]Edge, 0, g.size)
	for i := 1; i <= g.size; i++ {
		g.parents[i] = i
	}
	g.q = NewQueue()
	g.minWeight = 0
}

func (g *Graph) Kruskal() {
	g.sortEdges()
	g.buildMinSpanningTree()
	g.printMinSpanningTree()
}

func (g *Graph) sortEdges() {
	seen := make(map[[2]int]bool)
	for i := 1; i <= g.size; i++ {
		from := i
		for j := 0; j < len(g.matrix[i]); j++ {
			to := g.matrix[i][j].To
			key := [2]int{min(from, to), max(from, to)}
			if !seen[key] {
				edge := Edge{Begin: from, End: to, Weight: g.matrix[i][j].Weight}
				g.edges = append(g.edges, edge)
				seen[key] = true
			}
		}
	}
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].Weight < g.edges[j].Weight
	})
}

func (g *Graph) buildMinSpanningTree() {
	for _, edge := range g.edges {
		begin := g.getParent(edge.Begin)
		end := g.getParent(edge.End)
		if begin == end {
			continue
		}
		g.q.Enqueue(edge)
		g.minWeight += edge.Weight
		g.parents[end] = begin
	}
}

func (g *Graph) getParent(v int) int {
	if g.parents[v] != v {
		g.parents[v] = g.getParent(g.parents[v])
	}
	return g.parents[v]
}

func (g *Graph) printMinSpanningTree() {
	for !g.q.IsEmpty() {
		edge := g.q.Dequeue().(Edge)
		fmt.Printf("%d - %d: %d\n", edge.Begin, edge.End, edge.Weight)
	}
	fmt.Printf("Вес рёбер минимального остовного дерева: %d\n", g.minWeight)
}
