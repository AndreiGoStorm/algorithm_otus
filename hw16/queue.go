package hw16

type Queue struct {
	list *List
}

func (q *Queue) Enqueue(item interface{}) {
	q.list.PushBack(item)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	item := q.list.Front()
	q.list.Remove(item)
	return item.Value
}

func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}

func NewQueue() *Queue {
	return &Queue{&List{}}
}
