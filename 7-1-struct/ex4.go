package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}


func main1() {  
    emp := Employee{
        firstName: "aa",
        lastName:  "aaa",
    }
    fmt.Println("Employee", emp)	//Employee 5 {John Paul 0 0}
}

func main2() {  
    emp6 := Employee{"bb", "bbb", 55, 6000}
    fmt.Println("First Name:", emp6.firstName)
    fmt.Println("Last Name:", emp6.lastName)
    fmt.Println("Age:", emp6.age)
    fmt.Printf("Salary: $%d", emp6.salary)
}

func main() {  
    var emp7 Employee
    emp7.firstName = "cc"
    emp7.lastName = "ccc"
    fmt.Println("Employee 7:", emp7) //Employee 7: {Jack Adams 0 0}
}