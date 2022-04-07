package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	a := []int{10, 20, 30, 40, 50}

	//same-same : blank []
	fmt.Println("n:n ", a[1:1], a[2:2], a[3:3], a[4:4], a[5:5])

	//n to n+1 : single value a[n]
	fmt.Println("n:n+1 ", a[0:1], a[1:2], a[2:3], a[3:4], a[4:5])

	//n to m-1
	fmt.Println("2:5 ", a[2:5]) // 3 4 5
	fmt.Println("1:4 ", a[1:4]) // 2 3 4
	fmt.Println("2:3 ", a[2:3]) // 3

	fmt.Println("1st Element ", a[:1], a[0:1])
	fmt.Println("Last Element ", a[len(a)-1:], a[len(a)-1:len(a)])

	//len : m-n
	//cap : len(array) - n
	n, m := 1, 4
	b := a[n:m]
	fmt.Println("Length :", len(b), b, m-n)
	fmt.Println("Capacity :", cap(b), len(a)-n)

}

//https://go.dev/play/p/UsyPprdmfJg