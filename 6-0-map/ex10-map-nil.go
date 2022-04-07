package main

import (
	"fmt"
)

func main() {

	map_nil()
	map_not_nil()

}

func map_nil() {
	var map1 map[string]int

	/*
		//Invalid operation: map1 == map2 (map can only be compared to nil).
		map2 := map[string]int{"three" : 3}
		if map1 == map2 {
		}
	*/

	fmt.Println(map1) //map[]
	if map1 == nil {
		fmt.Println("map is empty") //map is empty
	}
}

func map_not_nil() {

	map1 := make(map[string]int)
	//map1 := map[string]int{}

	fmt.Println(map1) //map[]
	if map1 == nil {
		fmt.Println("map is empty") //not printed
	}

}

//https://go.dev/play/p/ZkieDV5sPNh