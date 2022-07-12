package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to  slices / vector boss , remember slice is another vector type ")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3]) // index 0 - 2 will be  print  rest will be cut out
	fmt.Println(fruitList)

	//fruitList = append(fruitList[kon index theke : kon index er age pojontho])

	fruitList = append(fruitList[1:]) // it will tream apple and print rest of slice
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777 error bcz we dont have 4th index

	// we can append it by this
	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	fmt.Println("this one set index 2 to empty")
	courses[2] = " "
	fmt.Println(courses)

	fmt.Println("this one will delete index 2 ")
	courses = append(courses[:index], courses[index+1:]...)
	//its work like [index , index)  0 1    ==> 2 skiped ==>   3 - n    pojontho print korbe .
	fmt.Println(courses)

}
