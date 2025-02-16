package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func ValidateStringLength(s any) (returnErr error) {
	returnErr = nil
	st := reflect.TypeOf(s)
	if st.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}
	sv := reflect.ValueOf(s)
	for i := 0; i < sv.NumField(); i++ {
		fieldVal := sv.Field(i)
		fieldType := st.Field(i)
		if fieldVal.Kind() != reflect.String {
			continue
		}
		lenS, ok := fieldType.Tag.Lookup("minStrlen")
		if !ok {
			continue
		}
		lenInt, err := strconv.ParseInt(lenS, 10, 64)
		if err != nil {
			returnErr = errors.Join(returnErr, err)
			continue
		}
		fvv := fieldVal.Interface().(string)	
		if len(fvv) < int(lenInt) {
			returnErr = errors.Join(returnErr,
				fmt.Errorf("length of string = %d is smaller than minStrlen = %d", int(lenInt), len(fvv)))
		}
	}

	return returnErr
}

type Turtle struct {
	A string
	B int `minStrlen:"15"`
	C string `minStrlen:"2"`
	D string `minStrlen:"3"`
	E string `minStrlen:"5"`
}

func main() {
	t := Turtle {
		A: "bananarama",
		B: 15,
		C: "C",
		D: "turtle",
		E: "ABCD",
	}	
	err := ValidateStringLength(t)
	fmt.Println(err)
}