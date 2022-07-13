package main

import "fmt"

type Person struct {
	name string
	age  int
}

//create function thn call struct
func NewPerson(pass_Name string) *Person {

	// age and name local cilo tobuo amra main block a eita use korty parbo karon amra pointer return korsi

	p := Person{name: pass_Name}
	p.age = 42
	p.name = "joy" // we dont use the passing name
	return &p
}

type User struct { // 1st latter is capital bcz user may needs to be exported out side of main
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	// This syntax creates a new struct.
	fmt.Println(Person{"Bob", 20})

	per := Person{"Joy", 22}
	fmt.Printf("created variable info is ==> %+v\n", per)

	// You can name the fields when initializing a struct.
	fmt.Println(Person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(Person{name: "Fred"})

	// An `&` prefix yields a pointer to the struct.
	fmt.Println(&Person{name: "Ann", age: 40})

	// It's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(NewPerson("Jon"))

	// Access struct fields with a dot.
	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	joy := User{"Joy", "joy@go.dev", true, 16} // create user joy
	fmt.Println(joy)
	fmt.Printf("details of joy are: %+v\n", joy) // %+v for detalied information
	fmt.Printf("Name is %v and email is %v.", joy.Name, joy.Email)

}
