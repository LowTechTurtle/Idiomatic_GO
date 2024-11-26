package main

import (
	"fmt"
)

func shadowing() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x, y := 5, 10
		fmt.Println(x, y)
	}
	fmt.Println(x)

	//this will not compile
	//	fmt := "banana"
	//	fmt.Println(fmt)
}
