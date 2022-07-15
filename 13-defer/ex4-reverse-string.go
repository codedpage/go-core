package main

import (
	"fmt"
)

func main() {
	name := "abcd"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}

//https://go.dev/play/p/PbTLnSDzueg