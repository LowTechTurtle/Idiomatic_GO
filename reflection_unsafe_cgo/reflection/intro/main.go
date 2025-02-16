package main

import (
	"fmt"
	"reflect"
)

type Turtle struct {
	Name string `turtle:"turtle_name"`
	Age  string `turtle:"turtle_age"`
}

func main() {
	var t []Turtle
	var tur Turtle
	tv := reflect.ValueOf(t)
	fmt.Println(tv.Type().Elem().Kind())
	fmt.Println(tv.Type())
	fmt.Println(tv.Kind())
	fmt.Println(tv.Type().Elem().Field(0))
	// fmt.Print(tv.Interface().([]Turtle))
	turv := reflect.ValueOf(tur)
	fieldKind := turv.Field(0).Kind()
	fmt.Println(fieldKind)
}
