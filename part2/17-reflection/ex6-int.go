package main

import (  
    "fmt"
    "reflect"
)

func main() {  
    a := 10
    x := reflect.ValueOf(a).Int()
    fmt.Printf("type:%T value:%v\n", x, x)
	
	
    b := "aa"
    y := reflect.ValueOf(b).String()
    fmt.Printf("type:%T value:%v\n", y, y)

}