package parser

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"
)

func Parse(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return nil, errors.New("empty")
	}
	count, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}
	if count > 2048 {
		return nil, errors.New("too big buffer")
	}
	if count < 0 {
		return nil, errors.New("negative number of lines")
	}
	out := make([]string, 0, count)
	for i := 0; i < count; i++ {
		if !scanner.Scan() {
			return nil, errors.New("too few lines")
		}
		text := scanner.Text()
		text = strings.TrimSpace(text)
		if len(text) == 0 {
			return nil, errors.New("empty line")
		}
		out = append(out, text)
	}
	return out, nil
}

func ToData(s []string) []byte {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(len(s)))
	b.WriteRune('\n')
	for _, v := range s {
		b.WriteString(v)
		b.WriteRune('\n')
	}
	return b.Bytes()
}
