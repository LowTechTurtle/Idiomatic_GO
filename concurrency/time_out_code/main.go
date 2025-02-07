package main

import (
	"context"
	"errors"
	"time"
)

func timeLimit[T any](worker func() T, limit time.Duration) (T, error) {
	out := make(chan T)
	ctx, cancel := context.WithTimeout(context.Background())

	defer cancel()
	go func() {
		out <- worker()
	}()

	select {
	case result := <-out:
		return result, nil
	case <-ctx.Done():
		var zero T
		return zero, errors.New("timed out")
	}
}
