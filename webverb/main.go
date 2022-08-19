package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	getpostreq()
}
func getpostreq() {
	const myurl = "http://localhost:3000/postform"
	data := url.Values{}
	data.Add("firstname", "joy adhikary")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	result, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(result))

}
