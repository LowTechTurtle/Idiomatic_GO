package main

// process func for business logic, its
// common to use closure to wrap business logic
func process(val int) int {
	return val + 1
}

func processConcurrently(inVals []int) []int {
	// create the channel
	in := make(chan int, 5)
	out := make(chan int, 5)
	// launch processing goroutines
	for i := 0; i < 5; i++ {
		go func() {
			for v := range in {
				out <- process(v)
			}
		}()
	}
	// load data in the in channel in another goroutine
	// read the data from the out channel
	var outVals []int = make([]int, 0, 5)
	for v := range out {
		outVals = append(outVals, v)
	}
	// return the data
	return outVals
}
