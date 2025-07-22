package hw10

import "fmt"

type BinaryTree struct {
	root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(key int) {
	node := NewNode(key)
	if bt.root == nil {
		bt.root = node
		return
	}

	bt.insertNode(node, bt.root)
}

func (bt *BinaryTree) insertNode(node, cur *Node) {
	if node.key < cur.key {
		if cur.left == nil {
			cur.left = node
			return
		}
		bt.insertNode(node, cur.left)
	} else {
		if cur.right == nil {
			cur.right = node
			return
		}
		bt.insertNode(node, cur.right)
	}
}

func (bt *BinaryTree) Sort(node *Node) string {
	if node == nil {
		return ""
	}
	s := bt.Sort(node.left)
	s += fmt.Sprintf("%d ", node.key)
	s += bt.Sort(node.right)
	return s
}

func (bt *BinaryTree) Search(key int) *Node {
	if bt.root == nil {
		return nil
	}
	return bt.SearchNode(key, bt.root)
}

func (bt *BinaryTree) SearchNode(key int, node *Node) *Node {
	if node == nil {
		return nil
	}

	if key == node.key {
		return node
	}
	if key < node.key {
		return bt.SearchNode(key, node.left)
	}
	return bt.SearchNode(key, node.right)
}

func (bt *BinaryTree) SearchParent(node, cur *Node) *Node {
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
		return bt.SearchParent(node, cur.left)
	}
	return bt.SearchParent(node, cur.right)
}

func (bt *BinaryTree) Remove(key int) {
	node := bt.Search(key)
	if node != nil {
		bt.RemoveNode(node)
	}
}

func (bt *BinaryTree) RemoveNode(node *Node) {
	parent := bt.SearchParent(node, bt.root)
	if node.right == nil {
		bt.noRightChildRemoving(parent, node)
		return
	}

	minimum, minParent := bt.searchMinimum(node)
	if parent != nil {
		if parent.left == node {
			parent.left = minimum
		} else {
			parent.right = minimum
		}
	} else {
		bt.root = minimum
	}

	if minimum != node.right {
		minParent.left = minimum.right
		minimum.right = node.right
	}
	minimum.left = node.left
}

func (bt *BinaryTree) noRightChildRemoving(parent, node *Node) {
	if parent != nil {
		if parent.left == node {
			parent.left = node.left
		} else {
			parent.right = node.left
		}
	} else {
		bt.root = node.left
	}
}

func (bt *BinaryTree) searchMinimum(node *Node) (*Node, *Node) {
	minimum := node.right
	minParent := node
	for minimum.left != nil {
		minParent = minimum
		minimum = minimum.left
	}
	return minimum, minParent
}
