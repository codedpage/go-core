package main

import (
	"fmt"
)

func main() {
	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}

// need in same line
//} else {

//https://go.dev/play/p/boa2gLVOtIN