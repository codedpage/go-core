package main

import (
	"fmt"
	"net/http"
)

func main() {

	//main1()
	//main2()
	//main3()
	main4()
}

//core - HandleFunc
func main1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HandleFunc")
	})
	http.ListenAndServe(":3001", nil)
}

//func - HandleFunc
func main2() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3002", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HandleFunc - function")
	//http.ServeFile(w, r, "index.hmt")
}

//core - Handle
func main3() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handle")
	}))
	http.ListenAndServe(":3003", nil)
}

//func - Handle
func main4() {
	http.Handle("/", Save())
	http.ListenAndServe(":3004", nil)
}

func Save() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handle - function")
	})
}
