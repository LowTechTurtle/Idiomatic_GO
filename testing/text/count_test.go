package count

import (
	"os"
	"testing"
)

func TestCountRune(t *testing.T) {
	f, err := os.Open("testdata/sample1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	n, err := count_rune(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	if n != 20 {
		t.Error("expected 20, got ", n)
	}
}