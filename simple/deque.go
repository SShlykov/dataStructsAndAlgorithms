package simple_structs

import (
	"container/list"
	"fmt"
)

type Deque struct {
	l *list.List
}

func NewDeque() *Deque {
	return &Deque{l: list.New()}
}

func (d *Deque) PushFront(value interface{}) {
	d.l.PushFront(value)
}

func (d *Deque) PushBack(value interface{}) {
	d.l.PushBack(value)
}

func (d *Deque) PopFront() interface{} {
	if d.l.Len() == 0 {
		return nil
	}
	return d.l.Remove(d.l.Front())
}

func (d *Deque) PopBack() interface{} {
	if d.l.Len() == 0 {
		return nil
	}
	return d.l.Remove(d.l.Back())
}

func Deq() {
	deque := NewDeque()
	deque.PushFront(1)
	deque.PushBack(2)
	fmt.Println("deque len: ", deque.l.Len())
}
