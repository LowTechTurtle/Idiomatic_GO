package main

func process(v int) int {
	// do some business logic
	return v + 123
}
func processChannel(ch chan int) []int {
	// already know exactly how many wanted to be launched
	const conc = 10
	results := make(chan int, conc)
	// doing this also helpful for preventing our services
	// from being overwhelmed
	for i := 0; i < 10; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	var out []int
	for i := 0; i < 10; i++ {
		out = append(out, <-results)
	}
	return out
}