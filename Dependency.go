package main

import "fmt"

// defineing structer

type str string

type intake struct {
	intake_info str
}

type student struct {
	Name   str
	ID     int64
	sec    str
	cgpa   float32
	intake intake
}

//function started

func get_string() str {

	fmt.Println("calling get_string function...( 1 )")
	return ("41")
}

func get_intake_info(m str) intake {

	fmt.Println("calling get_intake_info....( 2 )")
	return intake{intake_info: m} //push that string
}

func (m intake) info() str {

	fmt.Println("calling info methods.....( 3 )")
	return m.intake_info // fatch that string

}

func get_student(m intake) *student {

	return &student{
		Name:   "joy adhikary",
		ID:     18192103062,
		sec:    "seven",
		cgpa:   3.95,
		intake: m,
	}
}

func (s *student) print() {

	msg := s.intake.info() // intake struct theke intake_info ber korbe
	fmt.Println()
	fmt.Println("Name of student is :", s.Name)
	fmt.Println("Intake is : ", msg)
	fmt.Println("SEC is :", s.sec)
	fmt.Println("ID is : ", s.ID)
	fmt.Println("CGPA is : ", s.cgpa)

}

func main() {

	fmt.Println("Lets work with class dependency...")

	msg := get_string()            // getting string
	intake := get_intake_info(msg) // fatch intake info
	student := get_student(intake) // pass intake and fatch infromation about student

	student.print()
}

// dependency
// student depands on get_student ==> get_student depands on get_intake_info ==> get_intake_info depands on get_string
