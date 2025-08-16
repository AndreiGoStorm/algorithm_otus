package hw11

type Node struct {
	key    int
	height int
	left   *Node
	right  *Node
}

func NewNode(key int) *Node {
	return &Node{key: key}
}

func (n *Node) updateHeight() {
	n.height = 1 + n.MaxInt(n.left.getNodeHeight(), n.right.getNodeHeight())
}

func (n *Node) getNodeHeight() int {
	if n == nil {
		return -1
	}
	return n.height
}

func (n *Node) MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
