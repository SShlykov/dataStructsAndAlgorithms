package simple_structs

import "fmt"

func tuple() {
	// Определение кортежа с двумя элементами (int и string)
	var myTuple struct {
		a int
		b string
	}

	myTuple.a = 10
	myTuple.b = "Hello"

	fmt.Println(myTuple) // {10, Hello}
}
