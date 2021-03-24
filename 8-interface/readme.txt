package main 
  
import "fmt"
  
 
func main() { 
  
    var val1 interface{} = 6.5
    var val2 interface{} = "aa"      
      
    fmt.Println("Value: ", val1.(float32)) 
    fmt.Println("Value: ", val2.(string)) 
} 