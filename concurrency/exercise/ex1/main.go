package main

import (
	"fmt"
	"sync"
)

func twoWriteOneRead() {
	ch := make(chan int)
	var wg1 sync.WaitGroup
	wg1.Add(2)

	for i := 0; i < 2; i++ {
		go func() {
			defer wg1.Done()
			for j := 0; j < 10; j++ {
				ch <- j
			}
		}()
	}

	go func() {
		wg1.Wait()
		close(ch)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for val := range ch {
			fmt.Println(val)
		}
	}()

	wg2.Wait()
}

func main() {
	twoWriteOneRead()
}