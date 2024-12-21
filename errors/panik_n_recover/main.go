package main

import "fmt"

/*
	func div60(A []int) {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println(v)
			}
		}()
		for _, v := range A {
			fmt.Println(60 / v)
		}
	}

	func main() {
		div60([]int{1, 2, 3, 0, 5})
	}
*/
func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}
