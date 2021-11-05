package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

// world, one, two, 543210

// hello 543210 two one world

func myDefer() {
	for i := 0; i < 6; i++ {
		defer fmt.Print(i)
	}
}
