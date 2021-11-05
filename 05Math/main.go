package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to math in golang.")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4

	// fmt.Println("The sum is: ", myNumberOne+int(myNumberTwo))

	// Random from math
	// random number.
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// Random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
