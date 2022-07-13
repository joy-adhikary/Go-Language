package main

import "fmt"

func main() {

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days { // for (auto i: vector ) this kind of format
		fmt.Println(days[i])
	}

	for _, day := range days { // we can ignor index , print only value
		fmt.Printf("value is %v\n", day)
	}

	// continue goto break

	number := 1
	for number < 10 {

		if number == 4 {
			goto joy
		}
		if number == 5 {
			number++
			continue
		}
		fmt.Println("Value is: ", number)
		number++
	}

joy:
	fmt.Printf("Hey joy your number is %v\n", number)

}
