package main

import . "fmt"

type Describer interface {
	Describe()
}

func main() {
	var d1 Describer
	if d1 == nil {
		Printf("%T %v \n", d1, d1)
	}

	//d1.Describe() //panic: runtime error: invalid memory address or nil pointer dereference
}

//https://go.dev/play/p/7yYJw25S3hu