package main

import (
	"compress/gzip"
	"io"
	"os"
)

func countLetters(r io.Reader) (map[string]int, error) {
	buffer := make([]byte, 2048)
	out := map[string]int{}
	for {
		b, err := r.Read(buffer)
		// we should read first then check error
		// because there might be some data already read before an err occur
		for _, v := range buffer[:b] {
			out[string(v)]++
		}

		if err != nil {
			if err == io.EOF {
				return out, nil
			} else {
				var zero map[string]int
				return zero, err
			}
		}
	}
}

func buildGzipReader(name string) (*gzip.Reader, func(), error) {
	r, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}

	gzr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}

	return gzr, func() {
		gzr.Close()
		r.Close()
	}, nil
}

// because gzip.reader implement io.reader, we can write a func
// to count letter that read gzip file
func countLettersGzip(name string) (map[string]int, error) {
	gzr, f, err := buildGzipReader(name)
	defer f()
	if err != nil {
		return nil, err
	}
	m, err := countLetters(gzr)
	if err != nil {
		return nil, err
	}
	return m, nil
}
