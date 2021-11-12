package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var signals []string = []string{"test"}
var mut sync.Mutex

func main() {
	weblist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://facebook.com",
	}
	for _, site := range weblist {
		go GetStatusCode(site)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func GetStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("Status code %d for %s\n", res.StatusCode, endpoint)
	}
}
