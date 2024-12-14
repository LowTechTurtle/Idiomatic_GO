package main

import (
	"fmt"
)

type MyInt int

func main() {
	// simple interface demonstration
	var myStringer fmt.Stringer
	var myIncrementer Incrementer
	pointerCounter := &Counter{}
	valueCounter := Counter{}

	myStringer = pointerCounter
	myStringer = valueCounter
	myIncrementer = pointerCounter
	myIncrementer = &valueCounter
	fmt.Println(myStringer)
	fmt.Println(myIncrementer)
	myIncrementer.Increment()
	fmt.Println(myIncrementer)

	// interface and nil( interface is nil if its type and value is nil)
	var pCounter *Counter
	fmt.Println(pCounter == nil) // prints true
	var incrementer Incrementer
	fmt.Println(incrementer == nil) // prints true
	incrementer = pCounter
	fmt.Println(incrementer == nil) // prints false

	//type assertion
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)
}
