package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", 1001)
}

func doSomethingCool(ctx context.Context) {
	rID := ctx.Value("request-id")
	fmt.Println(rID)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println(rID)
}

func main() {
	fmt.Println("Go Context Tutorial")

	//signal breaker
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = enrichContext(ctx)
	go doSomethingCool(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
	}
	time.Sleep(2 * time.Second)
}
