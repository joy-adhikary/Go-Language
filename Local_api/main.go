package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
	Boss        *Boss   `json:"boss"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}
type Boss struct {
	Fullname string `json:"fullname"`
	Lastname string `json:"lastname"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h>hello</h>"))
}

func getallcourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("all the info is here ")
	w.Header().Set("content-Type ", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getonecourse id")
	w.Header().Set("content-Type ", "application/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("not found")
	return
}

func createonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	var course Course

	if r.Body == nil {
		json.NewEncoder(w).Encode("plz enter some data")
		return
	}

	//check doublicate
	for _, course1 := range courses {
		if course1.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("already exist")
			return
		}
	}

	json.NewDecoder(r.Body).Decode(&course)
	rand.Seed(time.Now().UnixNano())               ///create random number
	course.CourseId = strconv.Itoa(rand.Intn(100)) // convert int to string
	courses = append(courses, course)              // push it on courses
	json.NewEncoder(w).Encode(course)              // sending json file format
}

func updateonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses, course)
			json.NewEncoder(w).Encode("hey value is succesfully updated ")
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}
