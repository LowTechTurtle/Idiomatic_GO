package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

//go:embed help
var helpinfo embed.FS
func main() {
	if len(os.Args) == 1 {
		printFile()
		os.Exit(0);
	}

	data, err := helpinfo.ReadFile("help/" + os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(data)
	fmt.Println(string(data))
}

func printFile() {
	fs.WalkDir(helpinfo, "help", 
	func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			_, fileName, _ := strings.Cut(path, "/")
			fmt.Println(fileName)
		}
		return nil
	})
}