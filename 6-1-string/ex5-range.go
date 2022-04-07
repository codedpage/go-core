package main

import (
	"fmt"
)

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("Index %d, Char %c, Hex %x \n", index, rune, rune)
	}
}

func main() {
	name := "Señor"
	printCharsAndBytes(name)
}

//https://go.dev/play/p/HFXLC2rvZWc