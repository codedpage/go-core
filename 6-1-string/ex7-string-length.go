package main

import (
	"fmt"
	"unicode/utf8"
)

func length(s string) {
	fmt.Println(utf8.RuneCountInString(s))
}
func main() {

	word1 := "Señor"
	length(word1)           //5
	fmt.Println(len(word1)) //6

	fmt.Println("========")

	word2 := "Pets"
	length(word2)           //4
	fmt.Println(len(word2)) //4
}

//https://go.dev/play/p/u3Hs6FLzKrq