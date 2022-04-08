/*
Name convension
filename : xxxx_test.go
Function : TestXxxxxx

e.g. addition_test.go

//run test case
go test //one or more test case
go test a_test.go

*/
package main

import "testing"

func TestOnePlusOne(t *testing.T) {
	//onePlusOne := 1 + 1    //pass
	onePlusOne := 1 + 0 //fail
	if onePlusOne != 2 {
		t.Error("Expected 1 + 1 to equal 2, but got", onePlusOne)
	}
}

//https://go.dev/play/p/YTfeZ0e7fes
