package main

import (
	"fmt"
	"context"
	 
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", 1001)
}

func doSomethingCool(ctx context.Context) {
	rID := ctx.Value("request-id")
	fmt.Println(rID)
}

func main() {
	fmt.Println("Go Context Tutorial")
	ctx := context.Background()
	ctx = enrichContext(ctx)
	doSomethingCool(ctx)
}
