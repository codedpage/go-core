package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	example1()
}

func example1() {
	type Article struct {
		Name string `json:"fname"`
		AuthorName string
		draft      bool //not exported in json due to lowercase.
	}
	article := Article{
		Name:       "JSON in Go",
		AuthorName: "Mal Curtis",
		draft:      true,
	}
	
//data, err := json.Marshal(article)
data, err := json.MarshalIndent(article, "", "  ")


	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

