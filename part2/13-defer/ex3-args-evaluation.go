package main

import (  
    "fmt"
)

func printA(a int) {  
    fmt.Println("a=", a)
}
func main() {  
    a := 5
    defer printA(a)
    a = 10
    fmt.Println("Before deferred function call a=", a)

}