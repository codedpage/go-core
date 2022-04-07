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
	fmt.Println("length is", len(personSalary))

}
//https://go.dev/play/p/R6Io1bppxfB