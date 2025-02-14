package adder

import "testing"

func Test_add(t *testing.T) {
	result := add(1, 1)
	if result != 2 {
		t.Errorf("expected 2, got %d", result)
	}
}
