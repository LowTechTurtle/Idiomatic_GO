package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func add2(i int, j int) (int, error) { return i + j, nil }
func sub2(i int, j int) (int, error) { return i - j, nil }
func mul2(i int, j int) (int, error) { return i * j, nil }
func div2(i int, j int) (int, error) {
	if j == 0 {
		return -1, errors.New("division by zero")
	} else {
		return i / j, nil
	}
}

func clevermap2() {
	var opMap = map[string]func(int, int) (int, error){
		"+": add2,
		"-": sub2,
		"*": mul2,
		"/": div2,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"5", "*", "7"},
		{"1", "-", "3"},
		{"3", "/", "0"},
	}

	for _, sos := range expressions {
		if len(sos) != 3 {
			fmt.Println("invalid expression", expressions)
			continue
		}

		p1, err := strconv.Atoi(sos[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := sos[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("aint know this operator, this ops aint smooth")
			continue
		}

		p2, err := strconv.Atoi(sos[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		res, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(res)
	}
}

func fileLen(s string) (int, error) {
	var len int
	f, err := os.Open(s)
	if err != nil {
		return -1, err
	}

	defer f.Close()

	buffer := make([]byte, 2048)
	for {
		count, err := f.Read(buffer)
		len += count
		if err != nil {
			if err != io.EOF {
				return -1, err
			} else {
				return len, nil
			}
		}
	}
}

func prefixer(pre string) func(string) string {
	return func(suf string) string {
		return pre + " " + suf
	}
}
