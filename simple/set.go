package main

import (
	"fmt"
)

func setList() {
	set := make(map[int]bool)

	set[1] = true
	set[2] = true
	set[3] = true

	_, exists := set[2]
	if exists {
		fmt.Println("set: ", set, " 2 есть в множестве")
	}

	delete(set, 2)

	_, exists = set[2]
	if !exists {
		fmt.Println("set: ", set, " 2 удален из множества")
	}
}
