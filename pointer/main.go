package main

import (
	"fmt"
)

func main() {
	var f int = 20
	failedUpdate(&f)
	fmt.Println(f)
	successUpdate(&f)
	fmt.Println(f)

	p1 := MakePerson("Turtle", "Banana", 10)
	p2 := MakePerson("Banana", "Turtle", 12)
	fmt.Println(p1)
	fmt.Println(p2)

	SoS := []string{"Fool of a Took", "You shall not pass",
		"Keep your forked tongue behind your teeth"}

	fmt.Println(SoS)
	UpdateSlice(SoS, "Hope is rekindled")
	fmt.Println(SoS)
	GrowSlice(SoS, "Hope is rekindled")
	fmt.Println(SoS)

	TenMilPerson()
}
