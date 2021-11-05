package main

import "fmt"

func main() {
	fmt.Println("Welcome to study of arrays in GOLANG")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"

	fruitList[3] = "Peach"

	fmt.Println("FruitList is", fruitList)
	fmt.Println("Length of fruitList", len(fruitList))

	var vegetableList = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("Vegetable list is :", vegetableList)

}
