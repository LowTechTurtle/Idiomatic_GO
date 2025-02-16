package main

import (
	"fmt"
	"reflect"
	"time"
)

func AddTimer(f any) any {
	return reflect.MakeFunc(reflect.TypeOf(f),
		func(args []reflect.Value) (results []reflect.Value) {
			before := time.Now()
			results = reflect.ValueOf(f).Call(args)
			after := time.Now()
			processtime := after.Sub(before)
			fmt.Println(processtime)
			return results
		}).Interface()
}

func timeMe(a int) int {
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	return result
}
func main() {
	timed := AddTimer(timeMe).(func(int) int)
	fmt.Println(timed(2))
}
