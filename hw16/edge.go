package hw16

type Edge struct {
	Begin  int
	End    int
	Weight int
}

func NewEdge(begin, end, weight int) *Edge {
	return &Edge{begin, end, weight}
}
