/*Pointer receivers vs value receivers*/
package mymethod

import (  
    "fmt"
)

type Employee1 struct {  
    name string
    age  int
}

//value receiver  - temp changes
func (e Employee1) changeName(newName string) {  
    e.name = newName
	fmt.Println(e)
}

//pointer receiver - fixed changes
func (e *Employee1) changeAge(newAge int) {  
    e.age = newAge
	fmt.Println(e)
}

func Method_pointer_receiver_ex2() {  
	e := Employee1{
		name: "Ram",
		age:  50,
	}
	fmt.Println("\n+++  method_pointer_receiver_ex2  +++++")	 

	fmt.Printf(e.name) //Ram
	e.changeName("Shyam") //no impact 
	fmt.Printf(e.name) //Ram

	fmt.Printf("\nage: %d", e.age) //50 - before change
	(&e).changeAge(30) //method calling (by pointer)
	//e.changeAge(30)      //method calling  
	fmt.Printf("\nage: %d", e.age) //30 - after change
	
	//debug
	fmt.Println(e)
}
