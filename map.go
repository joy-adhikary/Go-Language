package main

import "fmt"

func main() {
	
	// sort using key  ... count the number of length(size) thn number thn  latter
	mp := make(map[string]string)
	mp["joy0"] = "adhikary"
	mp["joy1"] = "adhikary"
	mp["joy"] = "adhikary"
	mp["joy2"] = "adhikary"
	mp["joya"] = "adhikary"
	mp["joyaa"] = "adhikary"
	

	fmt.Println(mp)
	
	// output ==> map[joy:adhikary joy0:adhikary joy1:adhikary joy2:adhikary joya:adhikary joyaa:adhikary]


	languages := make(map[string]string) // map  of string-key  string-value
	// map[key]value

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Go"
	// insert in key-sorted  order

	fmt.Println("List of all languages: ", languages)

	fmt.Println("JS shorts for: ", languages["JS"]) //accessing index by key

	delete(languages, "RB") // delete any element by key
	fmt.Println("After removing ruby List of all languages: ", languages)

	languages["Abc"] = "BBBBb"

	fmt.Println("After insert Abc List of all languages: ", languages)

	// loops in map  in golang
	for key, value := range languages { // for key , value := range map_name
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	//if we wnat to ignor key then ,syntax=>  _,value := range map_name
	for _, value := range languages {
		fmt.Printf("value is %v\n", value)
	}

	//if we wnat to ignor value then ,syntax=>  key,_ := range map_name
	for key, _ := range languages {
		fmt.Printf("key is %v\n", key)
	}

}
