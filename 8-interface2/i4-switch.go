package main

import "fmt"

func findType(i interface{}) {

	fmt.Printf("\nType: %T \n", i)

	switch i.(type) {
	case string:
		fmt.Println("String: ", i.(string))

	case int:
		fmt.Println("Int:", i.(int))

	default:

		fmt.Println("undefined")
	}
}

func main() {

	findType("aaa")

	findType(10)

}

//https://go.dev/play/p/VYdQwMmRP9C
