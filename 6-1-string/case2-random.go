package main

import (
	"crypto/rand"
	"fmt"
)

func main() {

	str := "12345abcde"
	var strLen int = len(str)

	var len byte = byte(strLen) //type casting

	id := make([]byte, 5)
	fmt.Println(id)

	rand.Read(id)
	fmt.Println(id)
	fmt.Println()

	var subIndex byte
	for i, b := range id {

		subIndex = byte(b % len)
		fmt.Println(i, "------", b, "====", subIndex)

		id[i] = str[subIndex]
		fmt.Println(i, "......", id[i], "....", subIndex)
		fmt.Println()

	}

	fmt.Println(id)
	fmt.Println(string(id))
}

//https://go.dev/play/p/z5Hk3pj_bmP