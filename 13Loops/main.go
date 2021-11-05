package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang.")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Saturday", "Thursday"}

	fmt.Println(days)

	// PrintDaysUsingForI(days)
	// ForIndexDay(days)
	ForConditionLoop()
}

func PrintDaysUsingForI(days []string) {
	for i := 0; i < len(days); i++ {
		fmt.Printf("DAY => %v\n", days[i])
	}
}

func ForIndexDay(days []string) {
	for index, day := range days {
		fmt.Printf("Index is %v and Value is %v\n", index, day)
	}
}

func ForConditionLoop() {
	var rogueValue int = 0

	for rogueValue <= 10 {
		fmt.Println("ROGUE VAL : ", rogueValue)
		if rogueValue == 8 {
			goto exit_pgm
		}
		rogueValue += 1
	}

exit_pgm:
	fmt.Println("Exited..")

}
