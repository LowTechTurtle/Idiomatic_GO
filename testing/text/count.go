package count

import (
	"bytes"
	"io"
	"os"
	"unicode/utf8"
)

func count_rune(filename string) (int, error) {
	var count int
	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	buff := make([]byte, 2048)

	for {
		l, err := f.Read(buff)
		count += utf8.RuneCount(bytes.TrimSpace(buff[:l]))
		if err != nil {
			if err != io.EOF {
				return count, err
			} else {
				return count, nil
			}
		}
	}
}
