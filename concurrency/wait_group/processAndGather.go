package main

import "sync"

func processAndGather[T, R any](in <-chan T, processor func(T) R, num int) []R {
	out := make(chan R, num)
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		// we will create {num} number of thread to process the in channel
		// assume that the processor functin run relatively slow
		// then by doing this, we can run multiple processor at the same time
		// => we speed up the program
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var result []R
	for v := range out {
		result = append(result, v)
	}
	return result
}
