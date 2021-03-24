package main

import "fmt"

func main() {
	var num1 int = 400

	switch {
	case num1 < 200:
		fmt.Println("It's equal to 98")
	case num1 == 200:
		fmt.Println("It's equal to 100")
	case num1 > 200:
		fmt.Println("It's not equal to 98 or 100")
	}

}

//switch 语句像极了if
//任何支持进行相等判断的类型都可以作为测试表达式的条件，包括int， string， 指针等。
