package main

import "fmt"

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	t2Slice := make([]T2, 0, len(s))
	for _, v := range s {
		t2Slice = append(t2Slice, f(v))
	}
	return t2Slice
}

func Reduce[T1, T2 any](s []T1, init T2, f func(T2, T1) T2) T2 {
	res := init
	for _, v := range s {
		res = f(res, v)
	}
	return res
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	words := []string{"Potato", "Banana", "Paple"}

	filtered := Filter(words, func(s string) bool {
		if s == "Banana" {
			return true
		}
		return false
	})
	fmt.Println(filtered)

	reduced := Reduce(words, 0, func(l int, s string) int {
		return l + len(s)
	})
	fmt.Println(reduced)

	mapped := Map(words, func(s string) int {
		return len(s)
	})
	fmt.Println(mapped)
}
