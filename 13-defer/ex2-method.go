package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func main() {
	p := person{
		firstName: "aa",
		lastName:  "bb",
	}
	defer p.fullName() //without defer : aa bbWelcome

	fmt.Printf("Welcome ") //Welcome aa bb

}

//https://go.dev/play/p/ElJgos0HhlJ
