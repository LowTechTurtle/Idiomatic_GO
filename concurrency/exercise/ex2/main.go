package main

import (
	"fmt"
	"sync"
)

func twoRoutines() (chan int, chan int) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	wg1.Add(1)
	wg2.Add(1)

	go func() {
		defer wg1.Done()
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer wg2.Done()
		for i := 0; i < 10; i++ {
			ch2 <- i + 10
		}
	}()

	go func() {
		wg1.Wait()
		close(ch1)
	}()

	go func() {
		wg2.Wait()
		close(ch2)
	}()

	return ch1, ch2
}

func main() {
	ch1, ch2 := twoRoutines()
	for count := 0; count < 2; {
		select {
		case v1, ok := <-ch1:
			if !ok {
				ch1 = nil
				count++
				break
			}
			fmt.Println("from chan 1", v1)
		case v2, ok := <-ch2:
			if !ok {
				ch2 = nil
				count++
				break
			}
			fmt.Println("from chan 2", v2)
		}
	}
}
