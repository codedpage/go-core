package main

import "fmt"

func main() {

	var val1 interface{} = 6.5 //float64 defalut
	var val2 interface{} = "aa"

	fmt.Println("Value: ", val1.(float64))
	fmt.Println("Value: ", val2.(string))
}

//https://go.dev/play/p/8IMNuYpfU7w