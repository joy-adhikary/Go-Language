package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano()) // random number ganaretor
	diceNumber := rand.Intn(6) + 1   // between 1 - 6
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough
	case 4:
		fmt.Println("you can move to 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")
	}

	//condition in switch
	age := 45
	switch {
	case age < 18:
		fmt.Println("Kid")
	case age >= 18 && age < 40:
		fmt.Println("Young")
	default:
		fmt.Println("Old")
	}

	// switch with statements
	switch age := 29; {
	case age < 18:
		fmt.Println("Kid")
	case age >= 18 && age < 40:
		fmt.Println("Young")
	default:
		fmt.Println("Old")
	}
}
