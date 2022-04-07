package main

import (
	"fmt"
)

func main() {
	personSalary := map[string]int{
		"aa": 10,
		"bb": 20,
	}
	personSalary["cc"] = 30
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["dd"] = 40
	fmt.Println("Person salary changed", personSalary)

}

//https://go.dev/play/p/UsgogGXfQS_9