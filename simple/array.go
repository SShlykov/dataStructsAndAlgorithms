package simple_structs

import "fmt"

func Array() {
	var arr [5]int

	for i := 0; i < 5; i++ {
		arr[i] = i * i
	}

	fmt.Println("array: ", arr)
}
