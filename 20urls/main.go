package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "http://localhost:5000/api/v1/beneficiary/los?payment=das4e4c8"

func main() {
	fmt.Println("Welcome to URLs in golang.")
	fmt.Println(myUrl)

	// parsing
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())

	qparams := result.Query()

	for _, v := range qparams {
		fmt.Println(v)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=suyash",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)

}
