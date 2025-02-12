package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	var sum int
	var i int
	for {
		select {
		case <-ctx.Done():
			fmt.Println("sum: ", sum, "iterations: ", i, ctx.Err())
			return
		default:
			rando := rand.Int() % 100_000_000
			sum += rando
			i++
			if rando == 1234 {
				fmt.Println("sum: ", sum, "iterations: ", i, "got 1234")
				return
			}
		}
	}
}
