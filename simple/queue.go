package main

import "fmt"

type Queue []int

func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() int {
	if len(*q) == 0 {
		panic("Queue is empty!")
	}

	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func queue() {
	var q Queue

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println("queue: ", q)
	fmt.Println("queue: (dequeue) ", q, " val: ", q.Dequeue())
	q.Enqueue(2)
	fmt.Println("queue: (enqueue) ", q)
}
