package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers.")

	myNumber := 23
	var ptr = &myNumber // Refrence means &

	fmt.Println("Value of actual pointer is :=", ptr)
	fmt.Println("Value :=", *ptr) // (*) gives the value of that memory offset.

	// Operation using value
	*ptr = *ptr * 2
	fmt.Println("ptr aftermultiplying by two :=", *ptr)
}
