package main

import "fmt"

func add(x int, y int, z int) {
	result := x + y
	result = result * z
	fmt.Println(result)
}
func skill(x string, y string, z string) string {
	result := x + y + z
	return result
}
func add_return(x int, y int, z int) (int, int) { // num of return variable = num of return type
	result1 := x + y
	result2 := result1 * z
	return result1, result2 // return multiple value
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"

}

func main() {
	var s1, s2 string
	s1 = "C++"
	s2 = "GO"
	s3 := "C"
	fmt.Println(skill(s1, s2, s3))
	var x, y, z int
	x = 20
	y = 33
	z = 100
	add(x, y, z)

	a, b := add_return(x, y, z)
	fmt.Println(a, b)

	proRes, myMessage := proAdder(2, 5, 8, 7, 3)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro Message is: ", myMessage)

}

//https://www.youtube.com/watch?v=vdm04bVzkLg&t=727s
