// line by line execution hoi go te kintu defer use korle oita akdom last a hoi
// last in fast out manner
package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer() // last in thats why it will run fast

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// hello, 43210, two, One, world
