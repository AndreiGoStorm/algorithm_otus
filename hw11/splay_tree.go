package hw11

import "fmt"

type SplayTree struct {
	root *Node
}

func NewSplayTree() *SplayTree {
	return &SplayTree{}
}

func (st *SplayTree) Insert(key int) {
	cur := st.insertKey(key)
	parent := st.searchParent(cur, st.root)
	for parent != nil {
		grand := st.searchParent(parent, st.root)
		if parent.left != nil && parent.left == cur {
			node := st.smallRightRotation(parent)
			st.setGrand(grand, parent, node)
			parent = grand
			continue
		}
		if parent.right != nil && parent.right == cur {
			node := st.smallLeftRotation(parent)
			st.setGrand(grand, parent, node)
			parent = grand
			continue
		}

		panic("no children for parent")
	}
}

func (st *SplayTree) setGrand(grand *Node, parent *Node, node *Node) {
	if grand != nil {
		if grand.left == parent {
			grand.left = node
		} else {
			grand.right = node
		}
	}
	parent = grand
}

func (st *SplayTree) insertKey(key int) *Node {
	if st.root == nil {
		st.root = NewNode(key)
		return st.root
	}

	node := st.searchNode(key, st.root)
	if node != nil {
		return node
	}

	node = NewNode(key)
	st.insertNode(node, st.root)
	return node
}

func (st *SplayTree) insertNode(node, cur *Node) {
	if node.key < cur.key {
		if cur.left == nil {
			cur.left = node
			return
		}
		st.insertNode(node, cur.left)
	} else {
		if cur.right == nil {
			cur.right = node
			return
		}
		st.insertNode(node, cur.right)
	}
}

func (st *SplayTree) smallLeftRotation(a *Node) *Node {
	b := a.right
	a.right = b.left
	b.left = a
	return st.smallRotation(a, b)
}

func (st *SplayTree) smallRightRotation(a *Node) *Node {
	b := a.left
	a.left = b.right
	b.right = a
	return st.smallRotation(a, b)
}

func (st *SplayTree) smallRotation(a, b *Node) *Node {
	if st.root == a {
		st.root = b
	}
	st.updateHeight(a)
	st.updateHeight(b)
	return b
}

func (st *SplayTree) updateHeight(node *Node) {
	if node == nil {
		return
	}
	node.updateHeight()
}

func (st *SplayTree) Search(key int) *Node {
	if st.root == nil {
		return nil
	}
	return st.searchNode(key, st.root)
}

func (st *SplayTree) searchNode(key int, node *Node) *Node {
	if node == nil {
		return nil
	}

	if key == node.key {
		return node
	}
	if key < node.key {
		return st.searchNode(key, node.left)
	}
	return st.searchNode(key, node.right)
}

func (st *SplayTree) searchParent(node, cur *Node) *Node {
	if node == nil || cur == nil {
		return nil
	}
	if node.key == cur.key {
		return nil
	}
	if cur.left != nil {
		if node.key == cur.left.key {
			return cur
		}
	}
	if cur.right != nil {
		if node.key == cur.right.key {
			return cur
		}
	}
	if node.key < cur.key {
		return st.searchParent(node, cur.left)
	}
	return st.searchParent(node, cur.right)
}

func (st *SplayTree) Print(node *Node) {
	if node == nil {
		return
	}
	st.Print(node.left)
	fmt.Printf("%d ", node.key)
	st.Print(node.right)
}

func (st *SplayTree) Remove(key int) {
	node := st.Search(key)
	if node != nil {
		st.removeNode(node)
	}
}

func (st *SplayTree) removeNode(node *Node) {
	parent := st.searchParent(node, st.root)
	if node.right == nil {
		st.noRightChildRemoving(parent, node)
		return
	}

	minimum, minParent := st.searchMinimum(node)
	if parent != nil {
		if parent.left == node {
			parent.left = minimum
		} else {
			parent.right = minimum
		}
	} else {
		st.root = minimum
	}

	if minimum != node.right {
		minParent.left = minimum.right
		minimum.right = node.right
	}
	minimum.left = node.left
}

func (st *SplayTree) noRightChildRemoving(parent, node *Node) {
	if parent != nil {
		if parent.left == node {
			parent.left = node.left
		} else {
			parent.right = node.left
		}
	} else {
		st.root = node.left
	}
}

func (st *SplayTree) searchMinimum(node *Node) (*Node, *Node) {
	minimum := node.right
	minParent := node
	for minimum.left != nil {
		minParent = minimum
		minimum = minimum.left
	}
	return minimum, minParent
}
