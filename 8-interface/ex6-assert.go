package main

import (
	"fmt"
)

func assert(i interface{}) {
	//s := i.(int) //error
	s := i.(string)
	fmt.Println(s)
}
func main() {
	var s interface{} = "aa bb"
	assert(s)
}

//https://go.dev/play/p/PQlC24v7Fj5