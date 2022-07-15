package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	// create a random integer as the ID
	rand.Seed(time.Now().Unix())

	id := int64(rand.Int())
	operation1(id)
}

func operation1(id int64) {
	// do some work
	log.Println("operation1 for id:", id, " completed")
	operation2(id)
}

func operation2(id int64) {
	// do some work
	log.Println("operation2 for id:", id, " completed")
}
