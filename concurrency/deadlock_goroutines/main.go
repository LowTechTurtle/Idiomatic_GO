package main

import "fmt"

// this program will deadlock
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		inGoroutine := 1
		ch1 <- inGoroutine
		fromMain := <-ch2
		fmt.Println("goRoutines: ", inGoroutine, fromMain)
	}()

	inMain := 2
	ch2 <- inMain
	fromGoroutine := <-ch1
	fmt.Println("Main: ", fromGoroutine, inMain)
}
