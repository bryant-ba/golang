package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc")) //strings.Index 返回字符串在父字符串里的索引（位置）
	fmt.Printf("the position of the last instance of \"Hi\" is:   ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi")) //strings.LastIndex 返回字符串在父字符串里最后出现的位置索引
}

//当子符串是非ASCII编码的字符串时，可以用strings.IndexRune函数定位字符
