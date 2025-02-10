package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type timey struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int `json:"year"`
	Hour       int `json:"hour"`
	Minute     int `json:"minute"`
	Second     int `json:"second"`
}

func buildjson(t time.Time) string {
	structtime := timey{
		DayOfWeek: t.Weekday().String(),
		DayOfMonth: t.Day(),
		Month: t.Month().String(),
		Year: t.Year(),
		Hour: t.Hour(),
		Minute: t.Minute(),
		Second: t.Second(),
	}
	jsontime, err := json.Marshal(structtime)
	if err != nil {
		return time.RFC3339
	}
	return string(jsontime)
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		t := time.Now()
		var s string
		// if r.Header.Get("Accept") == "json" {
		// 	s = buildjson(t)
		// } else {
		// 	s = t.Format(time.RFC3339)
		// }
		s = buildjson(t)
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
