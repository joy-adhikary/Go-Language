package main

import "fmt"

const LoginToken string = "ghabbhhjd" // Public or global

func main() {
	var username string = "joy_adhikary"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username) // T for type

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 // upto 0 to 255  8 bit  2^8 digit
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable) // 0
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "github.com/joy-adhikary"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)
	// no var style

	numberOfUser := 300000.0 // another way to declare
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken) // using global variable
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
