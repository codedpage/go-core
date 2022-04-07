package main

import "fmt"

func main() {
	ch := make(chan string)
	select {
	case <-ch:
	default:
		fmt.Println("default case executed")
	}
}

//https://go.dev/play/p/8GN87oZLW7y