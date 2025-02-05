package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

func fileChecker(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in file checker: %w", err)
	}
	file.Close()
	return nil
}

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return slices.Equal(me.Codes, me2.Codes)
	}
	return false
}

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := re.Resource == other.Resource
		matchCode := re.Code == other.Code
		if ignoreResource {
			return matchCode
		} else if ignoreCode {
			return matchResource
		} else {
			return matchResource && matchCode
		}
	}
	return false
}

func main() {
	err := fileChecker("this_file_aint_exists.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("nigga, that file aint exists")
		}
	}

	maherrors := MyErr{Codes: []int{1, 2, 3, 4, 5}}

	if errors.Is(maherrors, MyErr{[]int{1, 2, 3, 4, 5}}) {
		fmt.Println("gud shit, nigga")
	}

	RE := ResourceErr{"Turtle", 1234}
	if errors.Is(RE, ResourceErr{Resource: "Turtle"}) {
		fmt.Println("Sneaky Turtle")
	}

	/*
	   err := AFunctionThatReturnsAnError()
	   var myErr MyErr
	   if errors.As(err, &myErr) {
	   fmt.Println(myErr.Codes)
	   }
	*/
}
