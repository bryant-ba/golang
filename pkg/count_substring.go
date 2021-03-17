package main

import (
	"fmt"
	"strings"
)

func main() {
	//var str string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggg"

	//fmt.Printf("Number of H's in %s is:  ",str)
	//fmt.Printf("%d\n",strings.Count(str,"H"))//strings.Count 函数返回字符串在父字符串中非重叠出现的次数

	fmt.Printf("Number of double g's in %s is:  ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "g"))
}
