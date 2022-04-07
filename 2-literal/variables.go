package main

import "fmt"
import "math"

func main() {
	var1()
	var2()
	var3()
	var4()
	var5()
	var6()
	var7()
	var8()
	var9()

}

func var1() {
	var age int // variable declaration
	fmt.Println("my age is", age)
}

func var2() {
	var age int // variable declaration
	fmt.Println("my age is ", age)
	age = 29 //assignment
	fmt.Println("my age is", age)
	age = 54 //assignment
	fmt.Println("my new age is", age)
}

func var3() {
	var age int = 29 // variable declaration with initial value

	fmt.Println("my age is", age)
}

func var4() {
	var age = 29 // type will be inferred

	fmt.Println("my age is", age)
}

func var5() {
	var width, height int = 100, 50 //declaring multiple variables

	fmt.Println("width is", width, "height is", height)
}

func var6() {
	var (
		name   = "naveen"
		age    = 29
		height int
	)
	fmt.Println("my name is", name, ", age is", age, "and height is", height)
}

func var7() {
	name, age := "naveen", 29 //short hand declaration

	fmt.Println("my name is", name, "age is", age)
}

func var8() {
	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)
}

func var9() {
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("minimum value is ", c)
}

//https://go.dev/play/p/mY7S7z9Yz_Y