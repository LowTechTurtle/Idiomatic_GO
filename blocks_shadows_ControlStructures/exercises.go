package main

import (
	"fmt"
	"math/rand"
)

func exercise() {
	var A []int = make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		A = append(A, rand.Intn(100))
	}

	for _, v := range A {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six")
		case v%2 == 0:
			fmt.Println("Two")
		case v%3 == 0:
			fmt.Println("Three")
		}
	}

	// in every loop, we shadow the variable total and gave it the value of the total of the outer scope
	// (whith is 0 because zero value of an int is 0) plus i. So its gonna print out 0 1 2 ... 9
	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}
