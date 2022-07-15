package main

import . "fmt"

type Employee struct {
	fn          string
	ln          string
	basicpay    int
	pf          int
	totalleaves int
	leavestaken int
}

func (e Employee) DisplaySalary() {
	Println(e.fn, e.ln, e.basicpay+e.pf)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalleaves - e.leavestaken
}

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

func main() {
	e := Employee{"aa", "bb", 5000, 200, 30, 5}

	var s SalaryCalculator = e
	s.DisplaySalary()

	var l LeaveCalculator = e
	Println(l.CalculateLeavesLeft())
}

//https://go.dev/play/p/iAlmi67wMyi
