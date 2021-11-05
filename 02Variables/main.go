package main

import "fmt"

// make first letter capital to declare public variable. 
const LoginToken string = "ansdhigafl.aksnduiasf.nasoiuyd9afhg.sdgoewftgrg"

func main() {
	var username string = "Suyashl13"
	fmt.Println(username)                              // Suyashl13
	fmt.Printf("Variable is of type: %T \n", username) // Variable is of type: string

	var isLoggedIn bool = true
	fmt.Printf("type of isLoggedIn => %T", isLoggedIn) // type of isLoggedIn => bool

	var smallVal uint16 = 256
	fmt.Println(smallVal)                           // 256
	fmt.Printf("typeof smallVal => %T\n", smallVal) // typeof smallVal => uint16

	var smallFloat float32 = 255.185768788474546
	fmt.Print(smallFloat) // 255.185780

	var anotherVariable int
	fmt.Println(anotherVariable)                                  // 0
	fmt.Printf("typeof anotherVariable => %T\n", anotherVariable) // typeof anotherVariable => int

	jwtToken := 3000000
	fmt.Printf("typeOf jwtToken => %T", jwtToken) // typeOf jwtToken => int
}
