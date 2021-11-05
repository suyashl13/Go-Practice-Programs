package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program."
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza : ")

	// commaok / err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating :", input)
}
