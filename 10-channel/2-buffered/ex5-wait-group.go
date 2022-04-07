/*
WaitGroup
The next section in this tutorial is about Worker Pools. To understand worker pools, we need to first know about WaitGroup as it will be used in the implementation of Worker pool.

A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is blocked until all Goroutines finish executing. Lets say we have 3 concurrently executing Goroutines spawned from the main Goroutine. The main Goroutines needs to wait for the 3 other Goroutines to finish before terminating. This can be accomplished using WaitGroup.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

//https://go.dev/play/p/4fKhrCYAM9k