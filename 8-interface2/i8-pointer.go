package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Desc() {
	fmt.Println(p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Desc() {

	fmt.Println(a.state, a.country)
}

type Describe interface {
	Desc()
}

func main() {

	//person
	var d1 Describe
	p := Person{"aa", 20}
	d1 = p //allowed
	//d1 = &p //allowed
	d1.Desc()

	//address
	var d2 Describe
	a := Address{"dl", "india"}
	//d2 = a //not allowed
	d2 = &a //allowed
	d2.Desc()
}

//https://go.dev/play/p/00MCsxfnJKZ
