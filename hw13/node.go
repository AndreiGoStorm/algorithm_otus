package hw13

const SIZE = 26

type Node struct {
	nodes []*Node
	value interface{}
}

func NewNode() *Node {
	return &Node{make([]*Node, SIZE), nil}
}

func (n *Node) next(c rune) *Node {
	index := n.getIndex(c)
	if !n.exists(c) {
		n.nodes[index] = NewNode()
	}
	return n.nodes[index]
}

func (n *Node) exists(c rune) bool {
	return n.nodes[n.getIndex(c)] != nil
}

func (n *Node) getIndex(c rune) rune {
	return c - 'a'
}
