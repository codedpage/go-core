package main

import (
	"fmt"
)

type Address struct {
	city, state string
}
type Contact struct {
	email, mobile string
}
type Person struct {
	name string
	age  int
	Address
	con []Contact
}

func main() {
	var p Person
	p.name = "aa"
	p.age = 20
	p.Address = Address{
		city:  "nd",
		state: "delhi",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)   //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field

	var p1 Person
	p1 = Person{
		name: "bb",
		age:  30,
		Address: Address{"nd",
			"delhi"},
		con: []Contact{Contact{"b@b.com", "101"}, Contact{"b1@b1.com", "102"}},
	}

	fmt.Println(p1)
}
