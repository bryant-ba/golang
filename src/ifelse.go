package main

import (
	"fmt"
)

func main()  {
	var (
		first,cond int
	)
	first = 10
	if first <= 0{

		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {

		fmt.Printf("first is between 0 and 5\n")
	} else {

		fmt.Printf("first is 5 or greeater\n")
	}
	if cond = 5; cond > 10 {

		fmt.Printf("cond is greater than 10\n")
	} else {

		fmt.Printf("cond is not greater than 10\n")
	}
}
