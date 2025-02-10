package main

import (
	"log/slog"
	"net/http"
	"time"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		t := time.Now()
		s := t.Format(time.RFC3339)
		_, err := w.Write([]byte(s))
		if err != nil {
			panic(err)
		}
	})

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			slog.Error("error serving", "Err", err)
		}
	}
}
