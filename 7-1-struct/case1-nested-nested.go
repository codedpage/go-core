package main

import (
	"fmt"
	_ "time"
)

var print = fmt.Println
var echo = fmt.Print
var printf = fmt.Printf
var br = fmt.Println

type address struct {
	city, state string
}

type contact struct {
	email, mobile string
}

type person struct {
	name string
	add  address
	con  []contact
}

func main() {

	var p person
	p.name = "aa"
	p.add = address{"nd", "delhi"}
	p.con = []contact{contact{"a@a.com", "101"}, contact{"a1@a1.com", "102"}}

	print(p)
	print(p.name, p.con)
	print(p.con[0].email)
	//----------------------------------------
	print("------------")
	var p1 person
	p1 = person{
		name: "bb",
		add:  address{"nd", "delhi"},
		con:  []contact{contact{"b@b.com", "101"}, contact{"b1@b1.com", "102"}},
	}

	print(p1)
	print(p1.name, p1.con)
	print(p1.con[0].email)
}
