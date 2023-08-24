package algorithms

import (
	"fmt"
)

func linearSearch(arr []int, x int) int {
	for i, item := range arr {
		if item == x {
			return i
		}
	}
	return -1
}

func LinearSearch() {
	arr := []int{10, 20, 80, 30, 60, 50, 110, 100, 130, 170}
	x := 110

	result := linearSearch(arr, x)
	if result != -1 {
		fmt.Printf("Элемент найден на позиции: %d\n", result)
	} else {
		fmt.Println("Элемент не найден в массиве.")
	}
}
