package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r)) // running it on server
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) { /// w for http reponsewriter  r for request
	// keu jodi respose send kore like perameter tahole eita r ar majhe tahkbe .. akn ei servehome fun jodi send respone  back kore tahole eita w te thakbe
	w.Write([]byte("<h1>Welcome to golang Module</h1>"))
}
