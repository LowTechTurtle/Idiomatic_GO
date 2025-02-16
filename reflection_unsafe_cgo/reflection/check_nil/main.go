package main

import (
	"fmt"
	"reflect"
)

func checkNil(v any) bool {
	vv := reflect.ValueOf(v)
	if !vv.IsValid() {
		return true
	}
	vk := vv.Kind()
	switch vk {
	case reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Map, reflect.Pointer, reflect.Slice:
		return vv.IsNil()
	default:
		return false
	}
}

func main() {
	var i interface{}
	var s string
	fmt.Println(checkNil(i))
	fmt.Println(checkNil(s))
}
