package main

import (
	"fmt"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func printCharsByte(s string) {
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%c ", bytes[i])
	}
}

func printCharsRune(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func main() {
	name := "Señor"

	printBytes(name)
	fmt.Println()
	printChars(name)
	fmt.Println()
	printCharsByte(name)
	fmt.Println()
	printCharsRune(name)
}

//https://go.dev/play/p/tTpAKg1aXTI