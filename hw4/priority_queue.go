package main

import "container/list"

type PriorityQueue struct {
	list *list.List
}

type QueueItem struct {
	priority int
	value    interface{}
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{list.New()}
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.list.Len() == 0
}

func (pq *PriorityQueue) getSize() int {
	return pq.list.Len()
}

func (pq *PriorityQueue) Enqueue(priority int, T interface{}) {
	newItem := &QueueItem{priority, T}
	for cur := pq.list.Front(); cur != nil; cur = cur.Next() {
		queueItem, _ := cur.Value.(*QueueItem)
		if priority < queueItem.priority {
			pq.list.InsertBefore(newItem, cur)
			return
		}
	}
	pq.list.PushBack(newItem)
}

func (pq *PriorityQueue) Dequeue() interface{} {
	if pq.IsEmpty() {
		return nil
	}

	item := pq.list.Front()
	pq.list.Remove(item)
	return item.Value
}
