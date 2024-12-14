package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *Person) String() string {
	if p == nil {
		return "Aint nobody here"
	}
	return fmt.Sprintf("%s %s, age: %d", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func (i Inner) ShowX() {
	fmt.Println(i.X)
}

type Inner struct {
	X int
}
type Outer struct {
	Inner
	X int
}
