package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// os.File implement both io.reader and io.writer
// so we can use os.File to demonstrate json encoder and decoder
func read_write() {
	toFile := Person{
		Name: "Fred",
		Age:  40,
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		panic(err)
	}

	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}

	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}

	var fromFile Person
	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", fromFile)
}
