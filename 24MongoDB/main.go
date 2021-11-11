package main

import (
	"fmt"
	"log"
	"mongodb-practice/router"
	"net/http"
)

// mongodb://localhost:27017/go_api

func main() {
	fmt.Println("Server is Getting started...")
	r := router.Router()

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
