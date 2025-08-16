package hw11

import (
	"fmt"
	"math/rand"
	"time"
)

type Treap struct {
	x     int
	y     int
	left  *Treap
	right *Treap
}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewTreap(x, y int, left, right *Treap) *Treap {
	if y <= 0 {
		y = rnd.Intn(100)
		//y = x + 1
	}
	return &Treap{x, y, left, right}
}

func CreateTreap(x int) *Treap {
	return NewTreap(x, 0, nil, nil)
}

func (t *Treap) Split(x int) (L *Treap, R *Treap) {
	if t == nil {
		return nil, nil
	}
	var newTree *Treap
	if t.x <= x {
		if t.right != nil {
			newTree, R = t.right.Split(x)
		}
		L = &Treap{
			x:     t.x,
			y:     t.y,
			left:  t.left,
			right: newTree,
		}
		return L, R
	} else {
		if t.left != nil {
			L, newTree = t.left.Split(x)
		}
		R = &Treap{
			x:     t.x,
			y:     t.y,
			left:  newTree,
			right: t.right,
		}
		return L, R
	}
}

func (t *Treap) Insert(x int) *Treap {
	l, r := t.Split(x)
	m := CreateTreap(x)
	return Merge(Merge(l, m), r)
}

func (t *Treap) Remove(x int) *Treap {
	l, r := t.Split(x - 1)
	_, r = t.Split(x)
	return Merge(l, r)
}

func (t *Treap) Search(key int) *Treap {
	if t == nil {
		return nil
	}
	return t.searchTreap(key, t)
}

func (t *Treap) searchTreap(key int, treap *Treap) *Treap {
	if treap == nil {
		return nil
	}

	if key == treap.x {
		return treap
	}
	if key < treap.x {
		return t.searchTreap(key, treap.left)
	}
	return t.searchTreap(key, treap.right)
}

func (t *Treap) Count(treap *Treap) int {
	return t.countNodes(treap, 0)
}

func (t *Treap) countNodes(treap *Treap, count int) int {
	if treap == nil {
		return count
	}
	count++
	count = t.countNodes(treap.left, count)
	count = t.countNodes(treap.right, count)
	return count
}

func Merge(l, r *Treap) *Treap {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.y > r.y {
		newRight := Merge(l.right, r)
		return NewTreap(l.x, l.y, l.left, newRight)
	} else {
		newLeft := Merge(l, r.left)
		return NewTreap(r.x, r.y, newLeft, r.right)
	}
}

func (t *Treap) Print(treap *Treap) {
	if treap == nil {
		return
	}
	t.Print(treap.left)
	fmt.Printf("[%d, %d]\n", treap.x, treap.y)
	t.Print(treap.right)
}
