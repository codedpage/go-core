package main

import (
	"fmt"
)

func main() {

	str := "12345abcde"

	for i, v := range str {

		fmt.Printf("%d -- %T -- %v == %c  # ", i, str[i], str[i], str[i])
		fmt.Println(v, string(v))

	}

}

/*
0 -- uint8 -- 49 == 1  # 49 1
1 -- uint8 -- 50 == 2  # 50 2
2 -- uint8 -- 51 == 3  # 51 3
3 -- uint8 -- 52 == 4  # 52 4
4 -- uint8 -- 53 == 5  # 53 5
5 -- uint8 -- 97 == a  # 97 a
6 -- uint8 -- 98 == b  # 98 b
7 -- uint8 -- 99 == c  # 99 c
8 -- uint8 -- 100 == d  # 100 d
9 -- uint8 -- 101 == e  # 101 e

*/

//https://go.dev/play/p/eyzTPWkZm_G
