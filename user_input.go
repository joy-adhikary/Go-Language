package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func fast_scan() {
	var n int
	fmt.Println("enter the vector size ==> ")
	fmt.Scan(&n) // taking input
	z := bufio.NewReader(os.Stdin)
	vector := make([]int, n)
	fmt.Println("enter the vector elimants ==> ")
	for i := 0; i < n; i++ {
		fmt.Fscan(z, &vector[i])
	}
	fmt.Println("given vector is ==> ", vector)
	sort.Ints(vector)
	fmt.Println("After sorting vector is ==> ", vector)
}
func main() {
	var joy int
	fmt.Println("enter the array size ==> ")
	fmt.Scan(&joy)
	array := make([]int, joy)
	fmt.Println("enter the array elimants ==> ")
	for i := 0; i < joy; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println("array is ", array)

	// another way to take input but a faster way
	fast_scan()
}
