package simple_structs

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{Data: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *LinkedList) Display() {
	current := l.Head
	fmt.Print("linked_list: [")
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Print("]\n")
}

func Linked_list() {
	list := &LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Display()
}
