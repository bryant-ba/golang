package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1) //An integer: 5, its location in memory: 0xc00000a088
}
