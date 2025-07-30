package hw16

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type List struct {
	Head  *ListItem
	Tail  *ListItem
	Count int
}

func (l *List) Len() int {
	return l.Count
}

func (l *List) Front() *ListItem {
	if l.Count == 0 {
		return nil
	}
	return l.Head
}

func (l *List) Back() *ListItem {
	if l.Count == 0 {
		return nil
	}
	return l.Tail
}

func (l *List) PushFront(v any) *ListItem {
	item := &ListItem{Value: v}
	if l.Head == nil {
		l.Head = item
		l.Tail = item
	} else {
		item.Next = l.Head
		l.Head.Prev = item
		l.Head = item
	}
	l.Count++
	return l.Head
}

func (l *List) PushBack(v any) *ListItem {
	item := &ListItem{Value: v}
	if l.Head == nil {
		l.Head = item
		l.Tail = item
	} else {
		item.Prev = l.Tail
		l.Tail.Next = item
		l.Tail = item
	}
	l.Count++
	return l.Tail
}

func (l *List) InsertAfter(v any, at *ListItem) *ListItem {
	item := &ListItem{Value: v, Next: at.Next, Prev: at}
	at.Next = item
	at.Next.Prev = item
	return item
}

func (l *List) InsertBefore(v any, at *ListItem) *ListItem {
	item := &ListItem{Value: v, Next: at, Prev: at.Prev}
	at.Prev = item
	at.Prev.Next = item
	return item
}

func (l *List) Remove(e *ListItem) {
	if e.Prev != nil {
		e.Prev.Next = e.Next
	} else {
		l.Head = e.Next
	}

	if e.Next != nil {
		e.Next.Prev = e.Prev
	} else {
		l.Tail = e.Prev
	}

	e.Next = nil
	e.Prev = nil
	l.Count--
}

func (l *List) MoveToFront(e *ListItem) {
	l.Remove(e)
	l.PushFront(e.Value)
}

func (l *List) insert(item, at *ListItem) *ListItem {
	item.Prev = at
	item.Next = at.Next
	item.Prev.Next = item
	item.Next.Prev = item
	l.Count++
	return item
}

func (l *List) Clear() {
	l.Head = nil
	l.Tail = nil
	l.Count = 0
}
