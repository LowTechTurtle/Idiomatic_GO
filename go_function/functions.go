package main

import (
	"errors"
	"fmt"
	"strconv"
)

type person struct {
	name   string
	age    int
	gender bool
}

func simulate_named_return_value(p person) {
	fmt.Println(p)
}

func variadic_arg_plus(x int, s ...int) (mapped []int) {
	/*
		mapped = make([]int, 0, len(s))
		for _, v := range s {
			mapped = append(mapped, v+x)
		}
	*/

	for i := 0; i < len(s); i++ {
		s[i] += x
	}
	return s
}

func div_and_remainder(num int, deno int) (int, int, error) {
	if deno == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / deno, num % deno, nil
}

// blank return
func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = num/denom, num%denom
	return
}

func f1(s string) int {
	return len(s)
}

func f2(s string) (r int) {
	for _, v := range s {
		r += int(v)
	}
	return
}

func f() {
	var funcvar func(s string) int
	funcvar = f1
	fmt.Println(funcvar("hello world"))
	funcvar = f2
	fmt.Println(funcvar("hello world"))
}

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

type opFuncType func(int, int) int

func clevermap() {
	//	var opMap = map[string]func(int, int) int{
	var opMap = map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"5", "*", "7"},
		{"1", "-", "3"},
		{"3", "/", "15"},
	}

	for _, sos := range expressions {
		if len(sos) != 3 {
			fmt.Println("invalid expression", expressions)
			continue
		}

		p1, err := strconv.Atoi(sos[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := sos[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("aint know this operator, this ops aint smooth")
			continue
		}

		p2, err := strconv.Atoi(sos[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		res := opFunc(p1, p2)
		fmt.Println(res)
	}
}

func anon_function() {
	for i := 0; i < 3; i++ {
		func(j int) {
			fmt.Println("Anoy function run number ", j)
		}(i)
	}
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}
