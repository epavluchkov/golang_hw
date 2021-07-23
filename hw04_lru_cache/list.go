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
	FirstItem *ListItem
	LastItem  *ListItem
	lenth     int
}

type ListImpl list

func (l *list) Len() int {
	return l.lenth
}

func (l *list) Front() *ListItem {
	return l.FirstItem
}

func (l *list) Back() *ListItem {
	return l.LastItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	li := ListItem{v, l.FirstItem, nil}

	if l.FirstItem != nil {
		l.FirstItem.Prev = &li
	}

	l.FirstItem = &li

	if l.LastItem == nil {
		l.LastItem = &li
	}

	l.lenth++

	return &li
}

func (l *list) PushBack(v interface{}) *ListItem {
	li := ListItem{v, nil, l.LastItem}

	if l.LastItem != nil {
		l.LastItem.Next = &li
	}

	l.LastItem = &li

	if l.FirstItem == nil {
		l.FirstItem = &li
	}

	l.lenth++

	return &li
}

func (l *list) Remove(item *ListItem) {
	if item.Prev != nil {
		item.Prev.Next = item.Next
	} else {
		l.FirstItem = item.Next
	}

	if item.Next != nil {
		item.Next.Prev = item.Prev
	} else {
		l.LastItem = item.Prev
	}

	l.lenth--
}

func (l *list) MoveToFront(item *ListItem) {
	if l.FirstItem != item {
		if item.Prev != nil {
			if l.LastItem == item {
				l.LastItem = item.Prev
			}
			item.Prev.Next = item.Next
		}

		if item.Next != nil {
			item.Next.Prev = item.Prev
		}

		l.FirstItem.Prev = item
		item.Prev = nil
		item.Next = l.FirstItem
		l.FirstItem = item
	}
}

func NewList() List {
	return new(list)
}
