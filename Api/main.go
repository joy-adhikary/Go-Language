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

// Model for course - file
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
	Lastname string `json:"Lastname"`
}

//making fake DB slice
var courses []Course

// middleware, helper  - file (check for unique coursename )
func (c *Course) IsEmpty() bool { // function IsEmpty for cheking passing course *

	// return c.CourseId == "" && c.CourseName == "" // if courseid and coursename is empty then return ture
	return c.CourseName == ""
}

func main() {
	fmt.Println(" Building an API ")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}, Boss: &Boss{Fullname: "Hitesh", Lastname: "Choudhary"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Joy Adhikary", Website: "go.dev"}, Boss: &Boss{Fullname: "JOY", Lastname: "Adhikary"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")                     // router r r.HandleFunc("rought", controller).Methods("methods name ")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")          // localhost:4000/courses dile ei rought ta show korbe
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")       // localhost:4000/courses/{id ta dibo kontar jonno dekhbo}
	r.HandleFunc("/course", createOneCourse).Methods("POST")        // POST is often used by World Wide Web to send user generated data to the web server or when you upload file.
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")    //PUT method is used to update resource available on the server
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE") // DELETE use for delete somethings form the db

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by JoyAdhikary</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json") // map type key value
	json.NewEncoder(w).Encode(courses)                 // it convert slice courses into json value
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course only")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	//params holo holder of  key value pair .... onk value thakte pare eitai
	params := mux.Vars(r) /// je id tar jonno request korbe sei id ta mux.vars(r) catch kore params ar moddhe rekhe dibe

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course) // if course id found then rap the value into json and through it
			return
		} else {
			json.NewEncoder(w).Encode("No Course found with given id") // if course id not found then rap the string into json and through it
			return
		}
	}

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data") // if data is nill
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // json.NewDecoder(kon value ke decode korbo).Decode(value ta kothai roise i mean kon slice or db a)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON
	for _, course1 := range courses {
		if course1.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("This course is already exist")
			return
		}
	}

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())               ///create random number
	course.CourseId = strconv.Itoa(rand.Intn(100)) // convert int to string
	courses = append(courses, course)              // push it on courses
	json.NewEncoder(w).Encode(course)              // sending json file format

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop => id => remove value =>  add new value with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course) // decode json values to normal value  from course
			course.CourseId = params["id"]              // assign id and update info
			courses = append(courses, course)           // push to slice
			json.NewEncoder(w).Encode("hey value is succesfully updated ")
			json.NewEncoder(w).Encode(course) // convert normal value to json value and send response
			return
		} else {
			json.NewEncoder(w).Encode("Course ID is not found in the database !!!")
		}
	}
	//TODO: send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Your course has been deleted!")
			break
		}
	}
}
