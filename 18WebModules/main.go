package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://lco.dev"

func main() {
	fmt.Println("Web requests.")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response := %T", response.Body)

	defer response.Body.Close() // callers responsibility to close connection.

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
