package main

import (
	"errors"
	"fmt"
)

// ~ before a type means it will work for the user defined concrete type 
// that have underlying type of the type specified in the Integer interface
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func DivAndRemainder[T Integer](num, deno T) (T, T, error) {
	if deno == 0 {
		return 0, 0, errors.New("denominator cant be 0")
	}
	return num / deno, num % deno, nil
}

type MahInt int

func main() {
	var a int = 1234231
	var b int = 1235
	var c uint32 = 123123
	var d uint32 = 321312
	var e MahInt = 321321
	var f MahInt = 321321

	fmt.Println(DivAndRemainder(a, b))
	fmt.Println(DivAndRemainder(c, d))
	fmt.Println(DivAndRemainder(e, f))
}