package main

func process(v int) {

}

// turn off a case using nil channel
func turnOffChan(in, in2 chan int) {
	for count := 0; count < 2; {
		select {
		case v, ok := <- in:
			if !ok {
				in = nil
				count++
			}
			process(v)
		case v, ok := <- in2:
			if !ok {
				in2 = nil
				count++
			}
			process(v)
		}
	}
}