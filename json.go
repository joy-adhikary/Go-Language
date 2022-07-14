package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // json:"coursename eita Name ke coursename a rename kore print korbe
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // json:"-"  "-" eita mane password ami show korbo nah .. sudu akta space dekhabe
	Tags     []string `json:"tags,omitempty"` //json:"tags,omitempty eitar omitempty mane jodi empty thake tahole dekhabo nah
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() { // amader kase je data ashe oigula json format a confert korbo

	lcoCourses := []course{
		{"ReactJS Bootcamp", 199, "joydhiakry.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 299, "joydhiakry.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 399, "joydhiakry.in", "hit123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // Indent onkta data ke part korar  mto kaj kore
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(` /// when data come from web it come in byte format 
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "joydhiakry.in",
		"tags": ["web-dev","js"]
	}
	`)
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb) // Valid is check jsonDataFromWeb is valid or not .. return bool type

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) //unmarshal take json data and convert it to row data and store it in the pointed variable (&lcocourse)
		fmt.Printf("%#v\n", lcoCourse)              // for printing interface we will use %#v
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{} // value can be any things that why we use interface insted of any fixed data type
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}
}
