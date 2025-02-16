package main

import (
	"fmt"
	"reflect"
)

func main() {
	st := reflect.TypeOf((*string)(nil)).Elem()
	sv := reflect.New(st).Elem()
	sv.SetString("banana")
	s := sv.Interface()
	fmt.Println(s)
	stringSliceType := reflect.TypeOf((*[]string)(nil)).Elem()
	ssv := reflect.MakeSlice(stringSliceType, 0, 10)
	ssv = reflect.Append(ssv, sv)
	ss := ssv.Interface().([]string)
	fmt.Println(ss, len(ss), cap(ss))
}