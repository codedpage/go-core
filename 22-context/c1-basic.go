package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
How to run

open two console
console - 1: For sever  :  server is ready

conosle - 2: For client :  hello controller, signal breaker (Ctl+c)
curl http://localhost:8091
ctl+c
Control goes here => case <-ctx.Done():

*/
func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8091", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "\nhello.....\n")
		fmt.Println("c1.....:")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		fmt.Println("c2.....:")
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
