package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang.")
	Greeter()

	var result int = adder(4, 5)

	fmt.Println("Result is :", result)
	fmt.Println("Pro Result is", proAdder(4, 875, 87, 1, 7, 2, 8, 54, 4))

	fs, ls := commaOkFunc("Hello", "World")
	fmt.Printf("FS : %v LS : %v", fs, ls)
}

func Greeter() {
	fmt.Println("Namaste from golang.")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	var total int = 0
	for _, val := range values {
		total += val
	}

	return total
}

func commaOkFunc(firstString string, lastString string) (string, string) {
	return firstString, lastString
}
