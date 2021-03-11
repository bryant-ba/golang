package main

import "fmt"

func main() {
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  //prints address Here is the pointer p: 0xc000088230
	fmt.Printf("Here is the string *p: %s\n", *p) //prints string  Here is the string *p: ciao
	fmt.Printf("Here is the string s: %s\n", s)   //prints same string Here is the string s: ciao
}
