package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var myTodos []Todo = []Todo{}

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	// IsDone bool   `json:"is_done"`
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myTodos)
}

func PostTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	for _, ktodo := range myTodos {
		if ktodo.Title == todo.Title {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/text")
			json.NewEncoder(w).Encode("Duplicate todo found.")
			return
		}
	}
	todo.Id = len(myTodos)
	myTodos = append(myTodos, todo)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myTodos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	for index, todo := range myTodos {
		if todo.Id == int(id) {
			myTodos = append(myTodos[:index], myTodos[index+1:]...)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(myTodos)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)

	badRequestJson := []byte(`
		{
			"err":"todo not found.",
		}
	`)

	json.NewEncoder(w).Encode(badRequestJson)
}

func main() {
	fmt.Println("TODO APP \nListening on Port no. 4000")
	router := mux.NewRouter()

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/", GetTodos).Methods("GET", "OPTIONS")
	router.HandleFunc("/", PostTodo).Methods("POST", "OPTIONS")
	router.HandleFunc("/{id}", DeleteTodo).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(":4000", handlers.CORS(header, methods, origins)(router)))
}
