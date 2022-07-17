package main

import (
	"fmt"
	"net/http"
)

 

func MyHandleFunction() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  
    fmt.Fprintf(w, "Hello Gopher through function...")
   
  })
}
 
 func main() {
	http.Handle("/", MyHandleFunction())
	http.ListenAndServe(":3000", nil)
}
