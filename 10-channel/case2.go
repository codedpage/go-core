package main

import (
	"fmt"

	"time"
)

var print = fmt.Println
var echo = fmt.Print
var printf = fmt.Printf
var br = fmt.Println

func main() {
	t := time.Now()

	n := 10

	for i := 1; i <= n; i++ {

		go process(i)
	}

	print("all-done")

	fmt.Println("Hello Go")
	print("Time Taken : ", time.Since(t))
}

func process(i int) {
	print(i)
	time.Sleep(2 * time.Second)

}

/*output
all-done
1
10
2
6
4
7
8
9
3
Hello Go
Time Taken :  976.1µs
5
*/

//https://go.dev/play/p/oXmCH4dHvQ2
