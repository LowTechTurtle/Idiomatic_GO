package main

import (
	"context"
	"net/http"
	"time"
)

func TimeOutContextFactory(milisec int) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancelfunc := context.WithTimeout(ctx, 
				time.Duration(milisec) * time.Millisecond)
			defer cancelfunc()
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}