package main

import (
	"fmt"
)

func main() {
	simulate_named_return_value(person{age: 20, name: "Turtle"})
	simulate_named_return_value(person{name: "Undying Turtle", gender: true})
	fmt.Println(variadic_arg_plus(1, 2, 3, 4, 2, 3, 1))

	_, _, err := div_and_remainder(5, 2)
	if err != nil {
		fmt.Println("Aint no error yet")
	}

	_, _, err = div_and_remainder(5, 0)
	if err != nil {
		fmt.Println(err)
	}

	x, _, err := divAndRemainder(5, 2)
	if err == nil {
		fmt.Println(x)
	}

	f()

	clevermap()
	anon_function()

	//shadowing some function in closure
	//closure can be passed in a function's return value
	//and they can access the base function's local val
	//essentially use 1 function's local var in another function! How fun.
	a := 20
	f := func() {
		fmt.Println(a)
		a := 30
		fmt.Println(a)
	}
	f()

	twoBase := makeMult(2)
	fmt.Println(twoBase(5))

	deferExample()

	//exercies
	clevermap2()
	fmt.Println(fileLen("a_test_of_ice_and_fire.txt"))

	hello := prefixer("Hello")
	bob := hello("Bob")
	fmt.Println(bob)
	maria := hello("Maria")
	fmt.Println(maria)
}
