package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 5
	var x64 int64 = int64(x)

	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(x64))

	fmt.Println(x, x64)

	var y interface{} = int64(x)
	var y64 int64 = y.(int64)

	fmt.Println(reflect.TypeOf(y))
	fmt.Println(reflect.TypeOf(y64))

	fmt.Println(y, y64)

}
