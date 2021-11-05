package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Study of If else in golang.")

	fmt.Printf("Enter marks in percentage : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	resultPercentage, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Please provide a valid int")
		os.Exit(1)
	} else if resultPercentage > 100 || resultPercentage < 0 {
		fmt.Println("Percentage must be between 0 - 100")
		os.Exit(1)
	}

	if resultPercentage < 35 {
		fmt.Println("Fail")
	} else if resultPercentage > 90 {
		fmt.Println("A+")
	} else if resultPercentage > 80 {
		fmt.Println("A")
	} else if resultPercentage > 70 {
		fmt.Println("B")
	} else if resultPercentage > 50 {
		fmt.Println("C")
	} else {
		fmt.Println("PASS Class")
	}

	ConditionalsDemo()

}

func ConditionalsDemo() {
	// Generate Random num
	rand.Seed(time.Now().UnixMicro())
	randomNum := rand.Intn(6 + 1)

	// Switch case demo
	switch randomNum {
	case 1:
		fmt.Println("Move One Spot")
	case 2:
		fmt.Println("Move Two Spots")
	case 3:
		fmt.Println("Move Three Spots")
	case 4:
		fmt.Println("Move Four Spots")
	case 5:
		fmt.Println("Move Five Spots")
	case 6:
		fmt.Println("Move Six Spots")
	}

}
