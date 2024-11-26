package main

import (
	"fmt"
)

func slices_declare_append() {
	var x []int
	x = append(x, 1, 2, 3, 4, 2, 3, 1)
	var y = []int{5, 6, 7}
	x = append(x, y...)
	fmt.Println(x)
}

func slices_capacity() {
	x := make([]int, 5)
	x = append(x, 1, 2, 3, 4, 5) //this will append 5 numbers after 5 0s
	fmt.Println(x)
	var y []int = make([]int, 0, 5)
	y = append(y, 1, 2, 3, 4, 5) // this will append 5 numbers and have len(y) = 5
	fmt.Println(y)
	clear(x) //clear the slice, but length is unchanged
	fmt.Println(x, len(x), cap(x))
}

func slice_slices() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	//there will be 2 zero string at the end after copying and if the length is 0 then no string copied(even if cap>0)
	var t []string = make([]string, 5)
	copy(t, x[1:])
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("t:", t)
}
