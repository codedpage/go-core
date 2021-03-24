package main

import (  
    "fmt"
)

func main() {  
	main0()
	main1()
	main2()
	
	//infinite loop
	//main3()
	
}

func main0() {  
    i := 0
    for ;i <= 10; { // initialisation and post are omitted
        fmt.Printf("%d ", i)
        i += 2
    }
}

func main1() {  
    i := 0
    for i <= 10 { //semicolons are ommitted and only condition is present
        fmt.Printf("%d ", i)
        i += 2
    }
}

func main2() {  
    for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
        fmt.Printf("%d * %d = %d\n", no, i, no*i)
    }

}


func main3() {  
    for {
        fmt.Println("Hello World")
    }
}