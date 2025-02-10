package main

import (
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"
)

func initLogger() *slog.Logger {
	options := slog.HandlerOptions{Level: slog.LevelInfo}
	handler := slog.NewJSONHandler(os.Stdout, &options)
	mySlog := slog.New(handler)
	return mySlog
}

var initLoggerCached func() *slog.Logger = sync.OnceValue(initLogger)

func logIp(ip string) {
	logger := initLoggerCached()
	logger.Info(ip)
}

func loggy(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logIp(r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

func main() {
	handler := http.NewServeMux()
	handler.Handle("/", loggy(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("yello"))
		})))

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