package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to study on slices in GOLANG.")

	var fruitList = []string{}
	fmt.Printf("Type of fruitList is %T \n", fruitList)
	fruitList = append(fruitList, "apples", "oranges", "mango", "carrots", "onions", "brocoli")

	fmt.Println("Fruit list :=", fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println("Fruit list :=", fruitList)

	highScore := make([]int, 4)
	highScore[0] = 534
	highScore[1] = 285
	highScore[2] = 284
	highScore[3] = 4744

	fmt.Println("HighScore := ", highScore)

	// Sort highScores array.
	sort.Ints(highScore)
	fmt.Println("HighScore := ", highScore)

	// Remove a value from slices based on index.
	var courses = []string{"react", "node", "django", "python", "ruby"}
	fmt.Println(courses)
	courses = append(courses[:2], courses[2+1:]...)
	fmt.Println(courses)
}
