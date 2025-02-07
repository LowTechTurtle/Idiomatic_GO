package main

import "sync"

type SlowComplicatedParser interface {
	Parse(string) string
}

func initParser() SlowComplicatedParser {
	// setup and things that is slow as shit
	var x SlowComplicatedParser
	// stub
	return x
}

var parser SlowComplicatedParser

var once sync.Once

func Parse(datatoParse string) string {
	once.Do(func() {
		parser = initParser()
	})
	return parser.Parse(datatoParse)
}

// using OnceValue
var initParserCached func() SlowComplicatedParser = sync.OnceValue(initParser)

func Parsev2(datatoParse string) string {
	parser = initParserCached()
	return parser.Parse(datatoParse)
}