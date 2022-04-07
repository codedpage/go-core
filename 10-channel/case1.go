package main

import (
	"fmt"
	"sync"
	"time"
)

var print = fmt.Println
var echo = fmt.Print
var printf = fmt.Printf
var br = fmt.Println

func main() {
	t := time.Now()

	n := 10
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {

		wg.Add(1)
		go process(i, &wg)
	}

	wg.Wait()
	print("all-done")

	fmt.Println("Hello Go")
	print("Time Taken : ", time.Since(t))
}

func process(i int, wg *sync.WaitGroup) {
	print(i)
	time.Sleep(2 * time.Second)
	wg.Done()
}

/*output
1
2
3
10
6
7
8
9
4
5
all-done
Hello Go
Time Taken :  2.0000953s
*/
//https://go.dev/play/p/ZsQFSGD5DoJ
