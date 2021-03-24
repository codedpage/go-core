package main

import (  
    "fmt"
)

//interface definition
type VowelsFinder interface {  
    FindVowels() []rune
}

type MyString string

//MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {  
    var vowels []rune
    for _, rune1 := range ms {
        if rune1 == 'a' || rune1 == 'e' || rune1 == 'i' || rune1 == 'o' || rune1 == 'u' {
            vowels = append(vowels, rune1)
        }
    }
    return vowels
}

func main() {  
    name := MyString("Sam Anderson")
    var v VowelsFinder
    v = name // possible since MyString implements VowelsFinder
    fmt.Printf("Vowels are %c", v.FindVowels())
	

	emp := map[string]interface{}{}
	emp["name"] = "rahul"
	emp["city"] = "delhi"
	emp["zip"] = 110056
	fmt.Println(emp)				//map[city:delhi name:rahul zip:110056]
	fmt.Printf("%T", emp)			//map[string]interface {}
	fmt.Println()
	dep := map[string]string{}
	dep["name"] = "rahul"
	dep["city"] = "delhi"
	dep["zip"] = "110056"
	fmt.Println(dep)				//map[city:delhi name:rahul zip:110056]
	fmt.Printf("%T", dep)			//map[string]string
	

}