package hw10

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

func (n *Node) getBalance() int {
	return n.left.getNodeHeight() - n.right.getNodeHeight()
}

func (n *Node) AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (n *Node) MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
