/*Pointer receivers vs value receivers*/
package mymethod

import (  
    "fmt"
)

type rectangle1 struct {  
    length int
    width  int
}

func perimeter(r *rectangle1) {   //function - pointer receiving (val not allowed)
    fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle1) perimeter() {  //method - pointer receiving (val can be allowed)
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func Method_pointer_receiver_ex1() { 
 
	fmt.Println("\n+++  method_pointer_receiver_ex1  +++++")	 
	//r-value, p-pointer
    r := rectangle1{
        length: 10,
        width:  5,
    }
    p := &r 
	
    perimeter(p)		//function - ref	- allowed
    //perimeter(r)		//function - val	- not allowed 
	
    p.perimeter()		//method - ref		- allowed
    r.perimeter()		//method - val		- allowed

}