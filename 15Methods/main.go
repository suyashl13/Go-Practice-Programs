package main

import "fmt"

func main() {
	suyash := User{"Suyash", "suyash.lawand@gmail.com", true, 12}
	suyash.GetStatus()
	fmt.Printf("Suyash %+v\n", suyash)
	suyash.NewMail()
	fmt.Printf("Suyash %+v\n", suyash)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("User Astive:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Printf("Suyash %+v\n", u)
}
