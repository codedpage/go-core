package main

import (
	"fmt"
)

type name struct {
	firstName string
	lastName  string
}

func main() {
	name1 := name{"aa", "xx"}
	name2 := name{"aa", "xx"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName: "bb", lastName: "yy"}
	name4 := name{}
	name4.firstName = "bb"
	//name4.lastName = "yy"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}

//https://go.dev/play/p/qlQBV-3EUTw