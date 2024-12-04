package main

import (
	"fmt"
	"time"
)

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func successUpdate(g *int) {
	*g = 15
}

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func MakePerson(FN string, LN string, age int) Person {
	return Person{FN, LN, age}
}

func MakePersonPointer(FN string, LN string, age int) *Person {
	p := Person{FN, LN, age}
	return &p
}

func UpdateSlice(SoS []string, S string) {
	SoS[len(SoS)-1] = S
	fmt.Println(SoS)
}

func GrowSlice(SoS []string, S string) {
	SoS = append(SoS, S)
	fmt.Println(SoS)
}

func TenMilPerson() {
	//var Persons []Person = make([]Person, 0, 5)
	var Persons []Person = make([]Person, 0, 10_000_000)
	start := time.Now()
	for i := 0; i < 10_000_000; i++ {
		Persons = append(Persons, Person{"Turtle", "Turlte", 20})
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
