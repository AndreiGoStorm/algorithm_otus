package huffman

type Node struct {
	value  byte
	weight int
	left   *Node
	right  *Node
}

func NewLeaf(value byte, weight int) *Node {
	return &Node{value: value, weight: weight}
}

func NewInternal(left, right *Node) *Node {
	return &Node{left: left, right: right, weight: left.weight + right.weight}
}
