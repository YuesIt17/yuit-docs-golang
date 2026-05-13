package hw04lrucache

// List — двусвязный список: все операции выполняются за O(1) без полного обхода списка.
type List interface {
	// Len возвращает количество элементов в списке.
	Len() int
	// Front возвращает первый элемент или nil, если список пуст.
	Front() *ListItem
	// Back возвращает последний элемент или nil, если список пуст.
	Back() *ListItem
	// PushFront вставляет значение v в начало списка и возвращает новый узел.
	PushFront(v interface{}) *ListItem
	// PushBack вставляет значение v в конец списка и возвращает новый узел.
	PushBack(v interface{}) *ListItem
	// Remove удаляет узел i из списка. i должен принадлежать этому списку.
	Remove(i *ListItem)
	// MoveToFront перемещает узел i в начало списка. i должен принадлежать этому списку.
	MoveToFront(i *ListItem)
}

// ListItem — элемент двусвязного списка.
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

// NewList возвращает пустой список.
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
