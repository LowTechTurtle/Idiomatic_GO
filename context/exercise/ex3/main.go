package main

import (
	"context"
	"fmt"
	"net/http"
)

type Level string

type contextkey int

const (
	_ contextkey = iota
	key
)

const (
	Info  Level = "info"
	Debug Level = "debug"
)

func LogFromContext(ctx context.Context) (Level, bool) {
	lvl, ok := ctx.Value(key).(Level)
	return lvl, ok
}

func ContextWithLog(ctx context.Context, log Level) context.Context {
	return context.WithValue(ctx, key, log)
}

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	// TODO get a logging level out of the context and assign it to inLevel
	inLevel, ok := LogFromContext(ctx)
	if !ok {
		return
	}
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

func MiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ll := r.URL.Query().Get("log_level")
		var lvl Level
		if ll == "info" {
			lvl = Info
		}
		if ll == "debug" {
			lvl = Debug
		}

		ctx := r.Context()
		ctx = ContextWithLog(ctx, lvl)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
