/*Value receivers in methods vs value arguments in functions

func:: sender:value - receiver:value
func:: sender:ref - receiver:pointer

func:: sender:value - receiver:pointer  X - not allowed
func:: sender:ref - receiver:value      X - not allowed

method:: sender:value/ref - receiver:value
method:: sender:value/ref - receiver:pointer

*/

package mymethod
import "fmt"

type rectangle struct {  
    length int
    width  int
}

//Receiver Function===========================

func area(r rectangle) {  
    fmt.Printf("Area - Function (sender:value, receive: value): %d\n", (r.length * r.width))
}

func area_func_pointer(r *rectangle) {  
  fmt.Printf("Area - Function (sender:ref, receive: pointer): %d\n", (r.length * r.width))
}

//Receiver Method==========================
func (r rectangle) area() {  
    fmt.Printf("Area + Method (sender:value/ref, receive: value): %d\n", (r.length * r.width))
}

func (r *rectangle) area_method_pointer() {  
    fmt.Printf("Area + Method (sender:value/ref, receive: pointer): %d\n", (r.length * r.width))
}

func Method_pointer_receiver() {  
	fmt.Println("\n+++  method_pointer_receiver  +++++")
    r := rectangle{
        length: 10,
        width:  5,
    }
    p := &r  

    //func:: sender:value - receiver:value====================================================
    area(r)      //area(*p), function call by value  
    //area(&r)   //area(p), function call by reference - not allowed

	//func:: sender:ref - receiver:pointer
    //area_func_pointer(r) //function call by value - not allowed
    area_func_pointer(p)   //function call by reference  
 
    fmt.Println("+++  ---------  +++++")
   
    //method:: sender:value/ref - receiver:value===================================================
    r.area()     //method call by value 
    (*p).area()

    p.area()    //method call by reference 
    (&r).area()

    fmt.Println("+++  ---------  +++++")
     
   //method:: sender:value/ref - receiver:pointer
    r.area_method_pointer()     //method call by value 
    (*p).area_method_pointer()

    p.area_method_pointer()    //method call by reference 
    (&r).area_method_pointer()
    
}
