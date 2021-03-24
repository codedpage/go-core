package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp8 := &Employee{"aa", "xx", 20, 2000}
    fmt.Println("First Name:", (*emp8).firstName) 
    fmt.Println("Age:", (*emp8).age)
	
    fmt.Println("First Name:", emp8.firstName)	//also allowed
    fmt.Println("Age:", emp8.age)				//also allowed
	
    emp81 := Employee{"bb", "xx", 40, 4000}	
    fmt.Println("First Name:", (emp81).firstName)
    fmt.Println("Age:", (emp81).age)	
}