
package main

import "fmt"

type Joy struct {
	name     string
	id       int
	position string
	salary   int
}

func (a Joy) show() { // func ( receiver_name (actually it is for catching purpose)   Data_type_of_receiver ) metod name ( para_meter ) return type { }

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Position Name: ", a.position)
	fmt.Println("ID articles: ", a.id)
	fmt.Println("Salary: ", a.salary)

}

func (a *Joy) show_pointer(x string) {

	a.position = x
}

type data int

func (x data) multi(y data) data {

	// method name multi, receiver x, parameter y, return type data

	return x * y

}

/*
if you try to run this code, then the compiler will throw an error
because  the type ( int ) and the method definitions have to present in the same package

func(d1 int)multi(d2 int)int{

        return d1 * d2
}

*/

func main() {

	res := Joy{
		name:     "joy",
		id:       203,
		position: "CSE",
		salary:   45000,
	}

	res.show() // receiver. method name diye call korty hoi

	v1 := data(10)
	v2 := data(30)

	ans := v1.multi(v2)

	fmt.Println()
	fmt.Println("Result of multi is : ", ans)

	// methods with pointer

	fmt.Println("Author's name: ", res.name)
	fmt.Println("Position Name(Before useing pointer ): ", res.position)

	p := &res // Creating a pointer
	p.show_pointer("ECE")

	fmt.Println()
	fmt.Println("Position Name(After useing pointer ): ", res.position)

	//  Go method can accept both value and pointer, whether it is defined with pointer or value receiver.. that means this also work ==>

	res.show_pointer("EEE")

	fmt.Println()
	fmt.Println("Position Name(After useing pointer ): ", res.position)

}




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







