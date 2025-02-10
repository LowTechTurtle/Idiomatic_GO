package main

import (
	"log/slog"
	"net/http"
	"time"
)

func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, r)
			end := time.Now()
			slog.Info("request time", "path", r.URL.Path,
			 "duration", end.Sub(start))
		})
}

var securityMsg = []byte("You didn't give the secret password\n")

func TerribleSecProvider(password string) func(http.Handler) http.Handler{
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("X-Secret-Password") != password {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write(securityMsg)
					return
				} else {
					h.ServeHTTP(w, r)
				}
			})
	}
}

func main() {
	terrbileSecProv := TerribleSecProvider("turtle")

	mux := http.NewServeMux()
	mux.Handle("/hello", terrbileSecProv(RequestTimer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello\n"))
		}))))
	
	// or to apply to every route in mux:
//	mux.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("Hello\n"))
//	}))
//	mux = terrbileSecProv(RequestTimer(mux))

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

}