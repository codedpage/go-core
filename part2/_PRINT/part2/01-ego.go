package main

import (
	"log"
	"net/http"
)

func ego() {
	// Simple static webserver:
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("D:/go_test/src/"))))
	mux.Handle("/gobyexample/", http.StripPrefix("/gobyexample", http.FileServer(http.Dir("D:/go_test/src/gobyexample/public/"))))

	log.Fatal(http.ListenAndServe(":3000", mux))
}
