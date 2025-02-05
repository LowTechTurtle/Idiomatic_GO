package main

import (
	"embed"
	"fmt"
)

//go:embed turtle
var noHidden embed.FS

//go:embed turtle/*
var parentHidden embed.FS

//go:embed all:turtle
var allHidden embed.FS

func main() {
	checkHidden("noHidden", noHidden)
	checkHidden("parentHidden", parentHidden)
	checkHidden("allHidden", allHidden)
}

func checkHidden(name string, dir embed.FS) {
	fmt.Println(name)
	allFileNames := []string{"turtle/.hiddenturtle", "turtle/secretturtle/.hiddenturtle"}
	for _, v := range allFileNames {
		_, err := dir.Open(v)
		if err == nil {
			fmt.Println(v, "found")
		}
	}
	fmt.Println()
}