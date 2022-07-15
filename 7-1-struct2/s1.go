package main

import "fmt"

type Employee struct {
	fn, ln      string
	age, salary int
}

func main() {

	e1 := Employee{"aa", "bb", 20, 5000}
	fmt.Println(e1)

	var e2 Employee
	fmt.Println(e2)

	e2 = Employee{
		fn:     "ak",
		ln:     "sharma",
		age:    30,
		salary: 8000,
	}
	fmt.Println(e2)
}

//https://go.dev/play/p/1zVItovWKN7
