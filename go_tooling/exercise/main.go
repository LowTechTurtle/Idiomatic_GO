package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed english_rights.txt
var eng_right string

//go:embed french_rights.txt
var fr_right string

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./binary language")
		os.Exit(1)
	}

	var lang_map map[string]string = map[string]string{
		"english": eng_right,
		"french":  fr_right}

	v, ok := lang_map[os.Args[1]]
	if !ok {
		fmt.Println("No rights for these mtf")
		os.Exit(1)
	}
	fmt.Println(v)
}
