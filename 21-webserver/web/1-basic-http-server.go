package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Gopher...")
	})

	http.ListenAndServe(":3000", nil)
}
 
 
 //using handle 
 func main_handle() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  
    fmt.Fprintf(w, "Hello Gopher through function...")
   
  }))
	http.ListenAndServe(":3000", nil)
}

 