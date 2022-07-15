package main

import "fmt"

type Address struct{ city, state string }
type Contact struct{ email, mobile string }
type Person struct {
	name    string
	age     int
	Address //promoted
}

func main() {
	var p Person
	p.name = "aa"
	p.age = 20

	//promoted
	p.Address = Address{
		city:  "ND",
		state: "Delhi",
	}

	fmt.Println(p.name, p.age)
	fmt.Println(p.city) //promoted
	fmt.Println(p)

	/////////////////////
	p1 := Person{
		name: "bb",
		age:  30,
		Address: Address{
			city:  "Patna",
			state: "Bihar",
		},
	}
	fmt.Println(p1.name, p1.age)
	fmt.Println(p1.city)
	fmt.Println(p1)

}

//https://go.dev/play/p/VDj7mQD6l2p
