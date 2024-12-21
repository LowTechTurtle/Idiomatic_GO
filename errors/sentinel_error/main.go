package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("This is not a zip file")
	notZip := bytes.NewReader(data)
	_, err := zip.NewReader(notZip, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Tola ya its not a zip file lul")
	}
}
