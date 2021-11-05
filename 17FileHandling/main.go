package main

import (
	"fmt"
	// "io"
	"io/ioutil"
	// "os"
)

func main() {
	fmt.Println("Welcome to file-handling in golang")

	// var content string = "This need to go in file - hiresuyash.com"

	// file, err := os.Create("./hiresuyashlog.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Length:", length)
	// defer file.Close()

	readFile("./hiresuyashlog.txt")

}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside file is \n", string(databyte))

}
