package huffman

import (
	"container/list"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Enqueue(node *Node) {
	for cur := q.list.Front(); cur != nil; cur = cur.Next() {
		n, _ := cur.Value.(*Node)
		if node.weight < n.weight {
			q.list.InsertBefore(node, cur)
			return
		}
	}
	q.list.PushBack(node)
}

func (q *Queue) Dequeue() *Node {
	if q.IsEmpty() {
		return nil
	}

	item := q.list.Front()
	q.list.Remove(item)
	return item.Value.(*Node)
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}
