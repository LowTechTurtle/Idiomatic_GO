package main

import (
	"fmt"
	"math/rand"
)

func if_and_for_statement() {
	// you can declare a variable that is local to the if block
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Old school for loop")
	}

	i := 0
	for ; i < 3; i++ {
		fmt.Println("Skip declare for loop")
	}

	for i < 5 {
		fmt.Println("Skip declare and increment for loop")
		i++
	}

	fmt.Println(i)
	/*
		for {
			fmt.Println("Infinite for loop")
		}
	*/
	// the for range loop
	s := []int{1, 2, 3}

	for i, v := range s {
		fmt.Println(i, v)
	}

	//skip index for
	for _, v := range s {
		fmt.Println(v)
	}

	//skip value for
	for i := range s {
		fmt.Println(i)
	}

	//for range in map
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k, v := range uniqueNames { //key, value in uniqueNames
		fmt.Println(k, v)
	}

	//for range used on string
	//for range iterate over rune, not bytes.
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	// for range value copy the value in the slice, not modifying them directly
	evenValues := []int{2, 4, 6, 8}
	for _, v := range evenValues {
		v += 1
		fmt.Println(v)
	}
	for _, v := range evenValues {
		fmt.Println(v)
	}

	// label for continue/break
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				break outer
			}
		}
		fmt.Println()
	}

}

func switch_statement() {
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}

	for _, v := range words {
		switch l := len(v); l {
		case 1, 2, 3, 4:
			fmt.Printf("%s is a short word\n", v)
		case 6, 7, 8, 9:
			fmt.Printf("%s is a long word\n", v)
		case 5:
			fmt.Printf("%s, perfectly balanced, as all things should be\n", v)
		default:
			fmt.Printf("%s, holy jesus, what heresy\n", v)
		}
	}

	// blank switches
	// in golang, switch is useful, it can be used instead of if statement
	// in fact we should do that if switch case increase the readability

	words = []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch {
		case len(word) < 5:
			fmt.Printf("%s is a short word\n", word)
		case len(word) == 5:
			fmt.Printf("%s, perfectly balanced, as all things should be\n", word)
		case len(word) > 5:
			fmt.Printf("%s is a long word\n", word)
		}
	}
}
