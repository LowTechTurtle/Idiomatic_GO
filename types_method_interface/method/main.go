package main

import "fmt"

const (
	Field1 = 0
	Field2 = 1 + iota
	Field3 = 20
	Field4
	Field5 = iota
)

type BitField int

const (
	bitField1 BitField = 1 << iota
	bitField2
	bitField3
	bitField4
)

func main() {
	turtle := Person{Age: 20, FirstName: "Almighty", LastName: "Turtle"}
	fmt.Println(turtle)
	var nilPerson *Person
	fmt.Println(nilPerson)

	// go automatically get the address and pass into the method
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	// if the type is pointer and the method takes value, go automatically
	// dereference it and pass into the method
	// this is just about syntatic sugar
	cp := &Counter{}
	fmt.Println(cp.String())
	cp.Increment()
	fmt.Println(cp.String())

	//lil fun iota
	fmt.Println(Field1, Field2, Field3, Field4, Field5)
	fmt.Println(bitField1, bitField2, bitField3, bitField4)

	// embedded field
	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)
	fmt.Println(o.Inner.X)
	o.ShowX()
}
