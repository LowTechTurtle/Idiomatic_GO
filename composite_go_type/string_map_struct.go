package main

import (
	"fmt"
)

func string_converting() {
	var a rune = 'x'
	var b byte = 'y'
	var x int = 72
	var s1 = string(a)
	var s2 = string(b)
	var s3 = string(x)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	var s string = "Bananarama"

	// if in a string have a character that UTF-8 cant encode, the rune slice will have an element that
	// fully represent that character, but the byte slice with break that into many byte => the length
	// of byte string >= the length of rune string when we convert the same string to 2 type of slice.
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}

func using_maps() {
	var nilMap map[string]int
	if nilMap == nil {
		fmt.Println("nilmap")
	}

	//map with 0 element but not nil
	var notNilMap map[string]int = map[string]int{}
	if !(notNilMap == nil) {
		fmt.Println("NotNilmap")
	}

	//if we know how much space in the map we're gonna use,
	// better use make with it to avoid dynamic allocate => faster execution
	teams := make(map[string][]string, 3)
	teams["Orcas"] = []string{"Fred", "Ralph", "Bijou"}
	teams["Lions"] = []string{"Sarah", "Peter", "Billie"}
	teams["Kittens"] = []string{"Waldo", "Raul", "Ze"}

	var totalWins = map[string]int{}
	totalWins["turtle"] = 1
	totalWins["rabbit"] = 2
	fmt.Println(totalWins["fish"])
	totalWins["fish"]++
	fmt.Println(totalWins["turtle"])
	fmt.Println(totalWins["rabbit"])
	fmt.Println(totalWins["fish"])
	fmt.Println(len(totalWins))
	delete(totalWins, "turtle")
	fmt.Println(len(totalWins))
}

func map_ok_idiom() {
	var m = map[string]int{}
	m["turtle"] = 10
	m["not turtle"] = 10
	_, ok := m["turtle"]
	if ok {
		fmt.Println("turtle is in the map")
	}

	_, ok = m["banana"]
	if !ok {
		fmt.Println("banana is not in the map")
	}

	//empty the map
	clear(m)
	fmt.Println(m, len(m))
}

func using_struct() {
	type person struct {
		name string
		age  int
	}

	var turtle person = person{"turtle", 200}
	fmt.Println(turtle)
	var ahoy person = person{age: 30, name: "ahoy"}
	fmt.Println(ahoy)

	//anonymous struct

	var banana struct {
		leng float32
		wid  float32
	}

	banana.leng = 3.22
	banana.wid = 0.55

	random_point := struct {
		x int
		y int
	}{
		x: 2,
		y: 3,
	}
	fmt.Println(random_point)
}
