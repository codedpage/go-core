package main

import (
	"fmt"
	"net/http"
)
/*
func main() {
	 
	mux := http.ServeMux{}
	mux.Handle("/users", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello users")
   
  })
	http.ListenAndServe(":3000", mux)
}
*/

func main() {
	 
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  
    fmt.Fprintf(w, "Hello mux server")
   
  }))
	http.ListenAndServe(":3000", mux)
}
