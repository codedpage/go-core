package main

import (
	"fmt"
)

func main() {
	//5 - Closures : a anonymous functions which can access the variables defined outside of function
	a5 := 5
	func() {
		fmt.Println("a5 =", a5)
	}()
}

//https://go.dev/play/p/2lyES6YVByh