package main

import (
	"fmt"
)

type Person struct {
	string
	int
}

func main() {
	main1()
	main2()
}

///////////////
func main1() {
	p := Person{"aa", 50}
	fmt.Println(p) //{aa 50}

	//not allowed - not in proper sequence
	//q := Person{ 40,"bb" }
	//fmt.Println(q)
}

func main2() {
	var p1 Person
	p1.string = "cc"
	p1.int = 20
	fmt.Println(p1) //{cc 20}
}

//https://go.dev/play/p/GS4puSk_ZiW
