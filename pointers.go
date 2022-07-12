package main

import "fmt"

func zeroval(ivalue int) {
	ivalue = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	fmt.Println("Welcome to  pointers in go ")

	var ptr2 *int
	fmt.Println("Value of pointer is ", ptr2)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)  // address
	fmt.Println("Value of actual pointer is ", *ptr) // address value

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)

	// why we use pointer here is the reason
	i := 1
	fmt.Println("initial:", i)

	zeroval(i) // increment but it work as local increments thats why i not increment
	fmt.Println("zeroval:", i)

	zeroptr(&i) // sending actual location /address thats why incerease on it mean increment on the actual value
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
