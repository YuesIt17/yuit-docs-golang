package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	back *ListItem
}

func NewList() List {
	return &list{
		head: &ListItem{},
	}
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head.Next
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	n := &ListItem{Value: v}
	n.Next = l.head.Next
	n.Prev = l.head
	if l.head.Next != nil {
		l.head.Next.Prev = n
	}
	l.head.Next = n
	if l.back == nil {
		l.back = n
	}
	l.len++
	return n
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.back == nil {
		return l.PushFront(v)
	}
	n := &ListItem{Value: v}
	n.Prev = l.back
	l.back.Next = n
	l.back = n
	l.len++
	return n
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if l.back == i {
		if i.Prev != l.head {
			l.back = i.Prev
		} else {
			l.back = nil
		}
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.head.Next == i {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if l.back == i {
		if i.Prev == l.head {
			l.back = nil
		} else {
			l.back = i.Prev
		}
	}
	i.Next = l.head.Next
	i.Prev = l.head
	if l.head.Next != nil {
		l.head.Next.Prev = i
	}
	l.head.Next = i
}
