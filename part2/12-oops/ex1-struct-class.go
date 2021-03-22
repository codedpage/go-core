package main

import (
	"fmt"
	"golanbot2/oops/employee"
)

func main() {
	ex1()
	fmt.Println()
	ex2()
}

func ex1() {
	e := employee.Employee{
		FirstName:   "Manoj",
		LastName:    "kumar",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}

func ex2() {
	//e := employee.Employee{"Manoj", "kumar", 30, 20}

	e := employee.New("Manoj", "kumar", 30, 20)
	e.LeavesRemaining()
}
