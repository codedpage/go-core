package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

//Is file exist
func main1() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)

		fmt.Println(err.Error())
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

//Is file exist
//https://golang.org/src/os/error.go?s=653:716#L11

func main2() {
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println(err) //err.Error()
		fmt.Println("File at path ==>", err.Path, "<== failed to open")

		fmt.Println(err.Op)
		fmt.Println(err.Path)
		fmt.Println(err.Err)

		fmt.Println(err.Error())
		fmt.Println(err.Unwrap())
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

//Is host exist
func main3() {
	addr, err := net.LookupHost("manojkumars.info")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

//IsExistFilePath
//https://golang.org/pkg/path/filepath/#Glob
//import "path/filepath"
func main4() {
	f, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		return
	}
	fmt.Println("matched files", f)
}

//supress error
func main() {
	f, _ := filepath.Glob("[")
	fmt.Println("matched files", f)
}

//https://go.dev/play/p/fx-sBT8xXU4