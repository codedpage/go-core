package main

import "fmt"
import "reflect"

func main() {

	/*
		var map1 map[string]int
			//Invalid operation: map1 == map2 (map can only be compared to nil).
			map2 := map[string]int{"three" : 3}
			if map1 == map2 {
			}
	*/

	map1 := make(map[string]int)
	map2 := make(map[string]int)

	map1["aaa"] = 10
	map2["aaa"] = 10

	eq := reflect.DeepEqual(map1, map2)
	if eq {
		fmt.Println("Two maps are equal")
	} else {
		fmt.Println("Two maps are not equal")
	}

}
