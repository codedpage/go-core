package main

import (
	"log"
	"math/rand"
	"time"
	"context"
)

// we need to set a key that tells us where the data is stored
const keyID = "id"

func main() {
	rand.Seed(time.Now().Unix())
	ctx := context.WithValue(context.Background(), keyID, rand.Int())
	log.Println("--1--", ctx)
	operation1(ctx)
}

func operation1(ctx context.Context) {
	// do some work

	// we can get the value from the context by passing in the key
	log.Println("operation1 for id:", ctx.Value(keyID), " completed")
	log.Println("--2--", ctx)
	operation2(ctx)
}

func operation2(ctx context.Context) {
	// do some work

	log.Println("--3--", ctx)
	// this way, the same ID is passed from one function call to the next
	log.Println("operation2 for id:", ctx.Value(keyID), " completed")
}