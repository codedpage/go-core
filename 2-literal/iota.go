package main

import (
	. "fmt"
)

func main() {

	//1
	const (
		a, b, c int = iota, iota, iota
	)
	Println(a, b, c)

	//2
	const (
		d int = iota
		e     //int = iota
		f     // int = iota
	)
	Println(d, e, f)

	//3
	const (
		w int = 10 + iota
		x     // int = iota
		y
		z
	)
	Println(w, x, y, z)

	//4
	const (
		i int = iota
		j
		k int = 50 + iota
		m
		n
	)
	Println(i, j, k, m, n)
}

//https://go.dev/play/p/zDztGN49cCE
