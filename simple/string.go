package main

import "fmt"

func myString() {
	var str1 string = "Hello"
	str2 := ", World!"

	message := str1 + str2
	length := len(message)
	char := message[1]
	fmt.Println("string: ", message, "; Длина строки:", length, "Второй символ:", char)
}
