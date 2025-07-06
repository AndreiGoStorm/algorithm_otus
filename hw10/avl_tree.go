package hw10

import (
	"fmt"
)

type AVLTree struct {
	root *Node
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (avl *AVLTree) Insert(key int) {
	node := NewNode(key)
	if avl.root == nil {
		avl.root = node
		return
	}

	avl.insertNode(node, avl.root)
}

func (avl *AVLTree) insertNode(node, cur *Node) {
	if node.key < cur.key {
		if cur.left == nil {
			cur.left = node
			return
		}
		avl.insertNode(node, cur.left)
	} else {
		if cur.right == nil {
			cur.right = node
			return
		}
		avl.insertNode(node, cur.right)
	}
	avl.updateHeight(cur.left)
	avl.updateHeight(cur.right)
	cur.left = avl.balance(cur.left)
	cur.right = avl.balance(cur.right)
}

func (avl *AVLTree) updateHeight(node *Node) {
	if node == nil {
		return
	}
	node.updateHeight()
}

func (avl *AVLTree) balance(node *Node) *Node {
	if node == nil {
		return nil
	}

	balance := node.getBalance()
	if node.AbsInt(balance) < 2 {
		return node
	}

	if balance > 1 {
		return avl.balanceLeftNode(node)
	}
	return avl.balanceRightNode(node)
}

func (avl *AVLTree) balanceLeftNode(node *Node) *Node {
	if (node.left.getBalance()) >= 0 {
		return avl.smallRightRotation(node)
	}
	return avl.bigRightRotation(node)
}

func (avl *AVLTree) balanceRightNode(node *Node) *Node {
	if (node.right.getBalance()) <= 0 {
		return avl.smallLeftRotation(node)
	}
	return avl.bigLeftRotation(node)
}

func (avl *AVLTree) smallLeftRotation(a *Node) *Node {
	b := a.right
	a.right = b.left
	b.left = a
	return avl.smallRotation(a, b)
}

func (avl *AVLTree) smallRightRotation(a *Node) *Node {
	b := a.left
	a.left = b.right
	b.right = a
	return avl.smallRotation(a, b)
}

func (avl *AVLTree) smallRotation(a, b *Node) *Node {
	if avl.root == a {
		avl.root = b
	}
	avl.updateHeight(a)
	avl.updateHeight(b)
	return b
}

func (avl *AVLTree) bigLeftRotation(a *Node) *Node {
	a.right = avl.smallRightRotation(a.right)
	a = avl.smallLeftRotation(a)
	return avl.bigRotation(a)
}

func (avl *AVLTree) bigRightRotation(a *Node) *Node {
	a.left = avl.smallLeftRotation(a.left)
	a = avl.smallRightRotation(a)
	return avl.bigRotation(a)
}

func (avl *AVLTree) bigRotation(a *Node) *Node {
	avl.updateHeight(a)
	return a
}

func (avl *AVLTree) Sort(node *Node) string {
	if node == nil {
		return ""
	}
	s := avl.Sort(node.left)
	s += fmt.Sprintf("%d ", node.key)
	s += avl.Sort(node.right)
	return s
}

func (avl *AVLTree) Search(key int) *Node {
	if avl.root == nil {
		return nil
	}
	return avl.SearchNode(key, avl.root)
}

func (avl *AVLTree) SearchNode(key int, node *Node) *Node {
	if node == nil {
		return nil
	}

	if key == node.key {
		return node
	}
	if key < node.key {
		return avl.SearchNode(key, node.left)
	}
	return avl.SearchNode(key, node.right)
}

func (avl *AVLTree) SearchParent(node, cur *Node) *Node {
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
		return avl.SearchParent(node, cur.left)
	}
	return avl.SearchParent(node, cur.right)
}

func (avl *AVLTree) Remove(key int) {
	node := avl.Search(key)
	if node != nil {
		avl.RemoveNode(node)
	}
}

func (avl *AVLTree) RemoveNode(node *Node) {
	parent := avl.SearchParent(node, avl.root)
	if node.right == nil {
		avl.noRightChildRemoving(parent, node)
		return
	}

	minimum, minParent := avl.searchMinimum(node)
	if parent != nil {
		if parent.left == node {
			parent.left = minimum
		} else {
			parent.right = minimum
		}
	} else {
		avl.root = minimum
	}

	if minimum != node.right {
		minParent.left = minimum.right
		minimum.right = node.right
	}
	minimum.left = node.left
	if node == minParent {
		node = minimum
	} else {
		node = minParent
	}

	for {
		if node == nil {
			break
		}
		avl.updateHeight(node)
		if node.left != nil {
			node.left = avl.balance(node.left)
		}
		if node.right != nil {
			node.right = avl.balance(node.right)
		}
		node = avl.SearchParent(node, avl.root)
	}
}

func (avl *AVLTree) noRightChildRemoving(parent, node *Node) {
	if parent != nil {
		if parent.left == node {
			parent.left = node.left
		} else {
			parent.right = node.left
		}
	} else {
		avl.root = node.left
	}
}

func (avl *AVLTree) searchMinimum(node *Node) (*Node, *Node) {
	minimum := node.right
	minParent := node
	for minimum.left != nil {
		minParent = minimum
		minimum = minimum.left
	}
	return minimum, minParent
}
