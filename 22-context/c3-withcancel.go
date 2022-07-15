package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

//Singal breaker (Cancel)
func operation1(ctx context.Context) error {
	// Let's assume that this operation failed for some reason
	// We use time.Sleep to simulate a resource intensive operation

	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
	//return nil
}

//main work
func operation2(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func main() {
	// Create a new context
	ctx := context.Background()
	fmt.Println(ctx)
	// Create a new context, with its cancellation function
	// from the original context
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println(ctx, "1----------", cancel)
	// Run two operations: one in a different go routine
	go func() {
		err := operation1(ctx)
		fmt.Println(ctx, "2============", err)
		// If this operation returns an error
		// cancel all operations using this context
		if err != nil {
			cancel()
		}
	}()

	// Run operation2 with the same context we use for operation1
	fmt.Println(ctx, "3============")
	operation2(ctx)
}
