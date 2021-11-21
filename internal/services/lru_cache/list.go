package lruCache

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
	Size int
	Head *ListItem
	Tail *ListItem
}

func (l list) Len() int {
	return l.Size
}

func (l list) Front() *ListItem {
	return l.Head
}

func (l list) Back() *ListItem {
	return l.Tail
}

func (l *list) PushFront(data interface{}) *ListItem {
	newFrontListItem := NewListItem(data)

	if l.Head == nil {
		l.Head, l.Tail = newFrontListItem, newFrontListItem
	} else {
		currentHead := l.Head
		newFrontListItem.Next, currentHead.Prev, l.Head = currentHead, newFrontListItem, newFrontListItem
	}
	l.Size++
	return newFrontListItem
}

func (l *list) PushBack(data interface{}) *ListItem {
	newBackListItem := NewListItem(data)
	currentTail := l.Tail
	if l.Tail == nil {
		l.Head, l.Tail = newBackListItem, newBackListItem
	} else {
		newBackListItem.Prev, currentTail.Next, l.Tail = currentTail, newBackListItem, newBackListItem
	}
	l.Size++
	return newBackListItem
}

func (l *list) Remove(item *ListItem) {
	switch p := item.Prev; {
	case p == nil:
		{
			if item.Next == nil {
				l.Head, l.Tail = nil, nil
			} else {
				l.Head = item.Next
				item.Next.Prev = nil
			}
		}
	case p != nil:
		{
			if item.Next == nil {
				l.Tail = item.Prev
				item.Prev.Next = nil
			} else {
				item.Prev.Next, item.Next.Prev = item.Next, item.Prev
			}
		}
		l.Size--
	}
}

func (l *list) MoveToFront(item *ListItem) {
	switch item {
	case l.Head:
		{
			break
		}
	case l.Tail:
		{
			l.Tail = item.Prev
			l.Tail.Next = nil
			currentHead := l.Head
			l.Head = item
			currentHead.Prev = item
			item.Next = currentHead
			item.Prev = nil
		}
	default:
		{
			currentHead := l.Head
			l.Head = item
			item.Next = currentHead
			currentHead.Prev = item
			item.Prev = nil
		}
	}
}

func NewList() List {
	return &list{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
}

func NewListItem(data interface{}) *ListItem {
	return &ListItem{
		Value: data,
		Next:  nil,
		Prev:  nil,
	}
}
