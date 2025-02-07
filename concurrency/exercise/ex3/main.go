package main

import (
	"fmt"
	"math"
	"sync"
)


func initSquareRoot() map[int]float64 {
	fmt.Println("called the slow baby")
	m := make(map[int]float64, 100_000)
	for i := 0; i < 10_000_000; i++ {
		m[i] = math.Sqrt(float64(i))
	}
	return m
}

var getmapsqrt func() map[int]float64 = sync.OnceValue(initSquareRoot)

func SquareRoot(val int) float64 {
	m := getmapsqrt()
	return m[val]
}

func main() {
	for i := 0; i < 10_000_000; i += 100000 {
		fmt.Println(SquareRoot(i))
	}
}