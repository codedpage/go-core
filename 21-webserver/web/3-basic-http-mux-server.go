package main

import (
	"fmt"
	"net/http"
)
 
func main0() {
	 
	mux := http.NewServeMux{}
	mux.Handle("/users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello users")
   
  }))
	http.ListenAndServe(":3001", mux)
 
}
 

func main() {
	 
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  
    fmt.Fprintf(w, "Hello mux server")
   
  }))
	http.ListenAndServe(":3001", mux)
}
