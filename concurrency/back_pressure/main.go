package main

import (
	"errors"
	"net/http"
	"time"
)

// system will perform better overall if their
// components limit the amount of work they will perform

type PressureGauge struct {
	// use empty struct to not waste memory
	// at this time of writing this comment, im not fully understand this
	// but this is kinda just use for a placeholder, like a semaphore maybe,
	// the idea is just throwing something that will not waste memory and
	// it will take a place, until a limit it hit and it wont accepting more work
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	// this channel with or without buffer still of type chan struct{}
	return &PressureGauge{
		ch: make(chan struct{}, limit),
	}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	// initialize an empty struct and write it to PG
	// if it did not hit the limit of how many goroutines
	// it should do at once, it will process
	case pg.ch <- struct{}{}: // this also occupy a place in the buffer
		// process
		f()
		// then read the channel => release the place in the buffer
		<-pg.ch
		return nil
	default:
		// return error if the limit of the PG is hit
		// this will assure that no more of {limit} gorountines running
		// at any point
		return errors.New("no more capacity")
	}
}

func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func main() {
	pg := New(10)

	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many request"))
		}
	})

	http.ListenAndServe(":8080", nil)
}
