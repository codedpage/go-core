package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "aaa": 10,
        "bbb": 20,
    }
    personSalary["ccc"] = 30
    fmt.Println("map before deletion", personSalary)
    delete(personSalary, "bbb")
    fmt.Println("map after deletion", personSalary)

}