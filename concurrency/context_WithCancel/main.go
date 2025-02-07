package main

import (
	"context"
	"fmt"
)

func countTo(ctx context.Context, max int) <-chan int {
	ch := make(chan int)
	go func() {
		// we need to defer ch here because we will quit this function
		// and the channel still waiting for someone to read the value
		// => the garbage collector will not collect this channel
		// so we need to defer close the channel to prevent goroutines leak
		defer close(ch)
		for i := 0; i < max; i++ {
			select {
			case <-ctx.Done():
				return 
			case ch <- i:
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// we need to defer cancel because when break is called and exited
	// the program exit without cancelling the context => we need to 
	// call cancel before exitting main
	defer cancel()
	ch := countTo(ctx, 10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}