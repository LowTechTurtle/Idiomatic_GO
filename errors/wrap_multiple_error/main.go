package main

import (
	"errors"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) ValidatePerson() error {
	var errs []error
	if len(p.FirstName) == 0 {
		errs = append(errs, errors.New("FirstName is empty"))
	}
	if len(p.LastName) == 0 {
		errs = append(errs, errors.New("LastName is empty"))
	}
	if p.Age < 0 {
		errs = append(errs, errors.New("Age cannot be negative"))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func main() {
	Anon := Person{"", "", -1}
	err := Anon.ValidatePerson()
	fmt.Println(err)
}
