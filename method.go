package main

import "fmt"

func main() {
	fmt.Println("methods in golang")

	Joy := User{"Joy", "Joy@go.dev", true, 16}
	fmt.Println(Joy)
	fmt.Printf("Joy details are: %+v\n", Joy)
	fmt.Printf("Name is %v and email is %v.\n", Joy.Name, Joy.Email)
	Joy.joys_methods()
	Joy.joy2_methods()
	fmt.Printf("Name is %v and email is %v.\n", Joy.Name, Joy.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (j User) joys_methods() { // passing user obj u  it can be anythings a b c ...z
	fmt.Println("Is user active: ", j.Name)
}

func (u User) joy2_methods() {
	u.Email = "test@go.dev" // it change locally ..
	fmt.Println("Email of this user is: ", u.Email)
}
