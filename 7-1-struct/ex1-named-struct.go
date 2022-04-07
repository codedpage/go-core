package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	//creating structure using field names
	emp1 := Employee{
		firstName: "aa",
		age:       25,
		salary:    500,
		lastName:  "aaa",
	}

	//creating structure without using field names
	emp2 := Employee{"bb", "bbb", 29, 800}

	fmt.Println("Employee 1", emp1) //Employee 1 {Sam Anderson 25 500}
	fmt.Println("Employee 2", emp2) //Employee 2 {Thomas Paul 29 800}

}

//https://go.dev/play/p/8OVi0LcK_1v
