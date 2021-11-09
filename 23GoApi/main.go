package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for Courses - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fakeDB
var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOne).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// routinh

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h2>Welcome to API by LCO</h2>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Look inside params
	for k, v := range params {
		fmt.Println(k, v)
	}

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("course not found")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course.")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
		return
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
	}

	// generate random unique id
	// append course into courses.

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET one course")
	w.Header().Set("Content-Type", "application/json")

	// first grab id from request
	params := mux.Vars(r)

	// loop, id, remove add with my id;
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
		}
		// TODO: send response when id is not found

	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

}
