package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("web verb lecture ")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() { // this one will create a get request to the server
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // closing request .. it will run at the end

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	// one way to print the response body
	// content, _ := ioutil.ReadAll(response.Body) // fatching all info from server body
	// fmt.Println(content)
	// fmt.Println(string(content)) // printing server  body info  by raping with string bcz we want to see in  string type

	//another way to print the response body
	// var responseString strings.Builder /// strings is more powerful package
	// content, _ := ioutil.ReadAll(response.Body)
	// byteCount, _ := responseString.Write(content)
	// fmt.Println("ByteCount is: ", byteCount)
	// fmt.Println(responseString.String())

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(` //create new rading format for json .. becoz we dont have any data in server side 
		{    // this one is json format 

			"coursename":"Let's go with golang",
			"price": 0,
			"platform":"JoyAdhikaryOnline.bd"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody) // creating post request // kon url a , kon format a  r kake post korbo 

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content)) // printing the post request body (application/json ) in string format
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "Joy")
	data.Add("lastname", "Adhikary")
	data.Add("email", "Joy@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
