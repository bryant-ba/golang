package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("test", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th")) //strings.HasPrefix 字符串的前缀
}
