package main

import "fmt"

func main() {
	fmt.Println("Maps in GOLANG")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)

	fmt.Println("Value of JS in map is =>", languages["JS"])

	for key, val := range languages {
		fmt.Printf("KEY : %v, VALUE : %v\n", key, val)
	}

}
