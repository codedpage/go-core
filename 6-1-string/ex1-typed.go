package main

import "fmt"

func main() {

	var a = "aaa" //string

	type myString string
	var b myString = "bbb" //myString

	//b = a		//not allowed because type does not matched (string, myStrin)

	//Re-assigning : allowed
    //a = "a1"
	//b = "b1"

	fmt.Printf("%v %T ", a, a)	//aaa string

	fmt.Printf("%v %T ", b, b)	//bbb main.myString
	
	c := 'u' //char 
	fmt.Printf("%v %T ", c, c)	//117 int32
}
