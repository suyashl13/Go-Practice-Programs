package main

func main() {
}

// Program to accept a number from command line and return prime or non prime number.
// func checkPrime(inputNum int) bool {
// 	var isPrime bool = true
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	inputNum, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
// 	checkNilAndExit(err)

// 	for i := 2; i < int(inputNum); i++ {
// 		if int(inputNum)%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	return isPrime

// }

// func checkNilAndExit(err error) {
// 	if err != nil {
// 		fmt.Println("Something went wrong..")
// 		os.Exit(1)
// 	}
// }

// check wheather there is a prime number in slice or not
// func checkContainsPrimenumber(numList []int) {
// 	for i, ele := range numList {
// 		var isPrime bool = checkPrime(ele)
// 		if isPrime {
// 			fmt.Printf("Prime number found at %v which is %v", i, ele)
// 			return
// 		}
// 	}
// }

// func appendTwoSlices() {
// 	sliceOne := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	sliceTwo := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

// 	sliceOne = append(sliceOne, sliceTwo...)

// 	fmt.Println(sliceOne)

// }

// func removeElementFromSlice() {
// 	var someSlice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	fmt.Println("Some slice", someSlice)
// 	// Remove 3
// 	someSlice = append(someSlice[:2], someSlice[3:]...)
// 	fmt.Println(someSlice)

// }
