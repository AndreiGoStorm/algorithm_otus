package hw13

type PrefixTree struct {
	root *Node
}

func NewPrefixTree() *PrefixTree {
	return &PrefixTree{NewNode()}
}

func (pt *PrefixTree) Insert(key string, value interface{}) {
	node := pt.root
	for _, ch := range key {
		node = node.next(ch)
	}
	node.value = value
}

func (pt *PrefixTree) Search(key string) interface{} {
	node := pt.getNode(key)
	if node != nil {
		return node.value
	}
	return nil
}

func (pt *PrefixTree) Delete(key string) {
	node := pt.getNode(key)
	if node != nil {
		node.value = nil
	}
}

func (pt *PrefixTree) getNode(key string) *Node {
	node := pt.root
	for _, ch := range key {
		if node.exists(ch) {
			node = node.next(ch)
		} else {
			return nil
		}
	}
	return node
}
