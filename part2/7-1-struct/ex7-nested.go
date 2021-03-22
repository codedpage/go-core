package main

import (
	"fmt"
)

type Address struct {
	city, state string
}

type Contact struct {
	mobile string
}

type Person struct {
	name    string
	age     int
	con     []Contact
	address Address
}

func main() {
	var p Person
	p.name = "aa"
	p.age = 20
	p.con = []Contact{Contact{"101"}, Contact{"102"}}
	//p.con = []Contact{Contact{mobile: "101"}, Contact{mobile: "102"}}
	p.address = Address{
		city:  "nd",
		state: "delhi",
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
	fmt.Println("Contact :", p.con)
	//////////////////////
	fmt.Println("------------")
	var p1 Person
	p1 = Person{
		name: "bb",
		age:  30,
		con: []Contact{Contact{"101"}, Contact{"102"}},
		address: Address{"nd",
			"delhi"},
	}

	fmt.Println(p1)
	fmt.Println(p.con[0].mobile)
	fmt.Println(p.address.state)
}
