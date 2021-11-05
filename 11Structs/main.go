package main

import "fmt"

func main() {
	suyash := User{"Suyash", "suyash@hiresuyash.com", true, 20}

	var john User = User{"John", "jks", true, 23}

	fmt.Printf("Johns detais : %+v\n", john)

	fmt.Printf("Suyash's details are := %+v\n", suyash)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

