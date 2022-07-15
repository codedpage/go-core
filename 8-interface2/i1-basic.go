package main

import "fmt"

func main() {

	//string
	var a interface{}
	a = "aa"
	fmt.Println(a.(string))
	fmt.Printf("a ======>%s %T \n", a, a)

	//int
	var b interface{}
	b = 20
	fmt.Println(b.(int))

	//string-1
	var s interface{} = "aa"
	fmt.Println(s.(string))

	//interface-array
	arr := [4]interface{}{1, "aa", 2.5, nil}
	fmt.Println("\ninterface-array")
	fmt.Println(arr)
	fmt.Println(arr[2])
	fmt.Printf("%T \n", arr)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	//interface-slice
	arr2 := []interface{}{1, "aa", 2.5, nil}
	fmt.Println("\ninterface-slice")
	fmt.Println(arr2)
	fmt.Println(arr2[2])
	fmt.Printf("%T \n", arr2)
	fmt.Println()

	//interface-in-function
	fmt.Println("\ninterface-in-function")
	show(arr2)
	fmt.Println()

}

func show(b []interface{}) {
	for i := 0; i < len(b); i++ {
		fmt.Print(b[i], " ")
	}
}

//https://go.dev/play/p/_czNXfY7nMZ
