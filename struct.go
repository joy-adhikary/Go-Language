package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	joy := User{"joy", "joy@go.dev", true, 16}
	fmt.Println(joy)
	fmt.Printf("joy details are: %+v\n", joy)
	fmt.Printf("Name is %v and email is %v.", joy.Name, joy.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
