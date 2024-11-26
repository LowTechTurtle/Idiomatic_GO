package main

import (
	"fmt"
)

const c int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

func main() {
	//float only save the closest value to the value specified
	var f float32 = 462356356.4568456123
	fmt.Printf("Random float: %f\n", f)

	// rune is just int32 but we shouldnt use int32 for rune
	var myFirstInitial rune = 'J' // good - the type name matches the usage
	var myLastInitial int32 = 'B' // bad - legal but confusing
	fmt.Println(myFirstInitial)
	fmt.Println(myLastInitial)

	//demonstrate explicit type conversion
	var a int32 = 5
	var b int64 = c
	d := int64(a) + b
	fmt.Println(d)

	//byte is uint8
	var x int = 100000
	var y byte = 100
	var sum3 int = x + int(y)
	var sum4 byte = byte(x) + y
	fmt.Println(sum3, sum4)

	//declaration list
	var (
		turtle        string = "turtle"
		num_of_turtle int    = 10
		turtlynormal  bool
	)

	fmt.Println(turtle, num_of_turtle, turtlynormal)

	//var overflow
	var realSmallI byte = 255
	var smallI int32 = 2147483647
	var bigI int64 = 9223372036854775807

	realSmallI += 1
	smallI += 1
	bigI += 1

	fmt.Println(realSmallI, smallI, bigI)
}
