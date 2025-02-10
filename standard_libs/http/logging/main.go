package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	//basic logging
	slog.Debug("Debug msg")
	slog.Info("Info msg")
	slog.Warn("Warn msg")
	slog.Error("Error msg")

	// can log with a message and then key value pair
	userId := "turtle"
	login_counts := 69420
	slog.Info("logging info", "turtle", userId, "login_counts", login_counts)

	// using json 
	options := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stderr, options)
	mySlog := slog.New(handler)

	turtleval := "turtling around"
	mySlog.Debug("turtle msg", "turtlekey", turtleval, "ID", userId)

	// faster loggin using logattrs
	mySlog.LogAttrs(context.Background(), slog.LevelDebug,
	 "fast turtle", slog.String("turtlekey", turtleval),
	slog.String("ID", userId))

	//using structured log to write old log
	myLog := slog.NewLogLogger(mySlog.Handler(), slog.LevelDebug)
	myLog.Println("using the mySlog Handler")
}