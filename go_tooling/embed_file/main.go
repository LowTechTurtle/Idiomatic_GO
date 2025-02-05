package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed password.txt
var passwords string

func main() {
	passwords_slice := strings.Split(passwords, "\n")
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			for _, u := range passwords_slice {
				if v == u {
					fmt.Printf("%s is a commonly used password\n", u)
				}
			}
		}
	} else {
		fmt.Println("Usage: ./'binary_name' password1 password2...")
	}
}
