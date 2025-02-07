package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		doThing1()
	}()

	go func() {
		defer wg.Done()
		doThing2()
	}()

	go func() {
		defer wg.Done()
		doThing3()
	}()

	wg.Wait()
}