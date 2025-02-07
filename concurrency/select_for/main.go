package main

import "fmt"

//solution to the deadlock scenario in deadlock_goroutines dir

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go func() {
		inGoroutines := 1
		ch1 <- inGoroutines
		fromMain := <- ch2
		fmt.Println("In GoRoutines: ", inGoroutines, fromMain)
	}()

	inMain := 2
	var fromGoroutine int
	select {
	case ch2 <- inMain:
	case fromGoroutine = <- ch1:
	}

	fmt.Println("In Main: ", fromGoroutine, inMain)
}