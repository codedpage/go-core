package main

import "fmt"

func main() {
	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")

	}
}

//https://go.dev/play/p/KZEeYuVE_9X