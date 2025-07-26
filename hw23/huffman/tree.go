package huffman

import "sort"

type Tree struct {
	q            *Queue
	root         *Node
	frequencyLen int
}

func NewTree() *Tree {
	return &Tree{q: NewQueue()}
}

func (t *Tree) BuildHuffmanTree(frequency map[byte]int) {
	t.frequencyLen = len(frequency)
	for _, node := range t.createNodes(frequency) {
		t.q.Enqueue(node)
	}

	for t.q.Len() > 1 {
		left := t.q.Dequeue()
		right := t.q.Dequeue()
		t.q.Enqueue(NewInternal(left, right))
	}
	t.root = t.q.Dequeue()
}

func (t *Tree) createNodes(frequency map[byte]int) []*Node {
	nodes := make([]*Node, 0, t.frequencyLen)
	for value, weight := range frequency {
		nodes = append(nodes, NewLeaf(value, weight))
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].value < nodes[j].value
	})
	return nodes
}

func (t *Tree) BuildCodeTable(root *Node) map[byte]string {
	codeTable := make(map[byte]string, t.frequencyLen)
	var getNode func(node *Node, code string)
	getNode = func(node *Node, code string) {
		if node.left == nil && node.right == nil {
			codeTable[node.value] = code
			return
		}
		if node.left != nil {
			getNode(node.left, code+"0")
		}
		if node.right != nil {
			getNode(node.right, code+"1")
		}
	}
	getNode(root, "")
	return codeTable
}
