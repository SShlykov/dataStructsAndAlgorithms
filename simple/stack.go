package simple_structs

import "fmt"

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	l := len(*s)
	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res
}

func Stack() {
	var s Stack

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println("stack: ", s)
	fmt.Println("stack: (pop)", s, "; pop_val: ", s.Pop())
	s.Push(569)
	fmt.Println("stack: (put)", s)
}
