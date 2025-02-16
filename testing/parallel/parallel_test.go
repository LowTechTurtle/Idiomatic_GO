package parallel

import "testing"

func TestParallel(t *testing.T) {
	data := []struct {
		name   string
		in     int
		out    int
		errMsg string
	}{
		{"a", 10, 20, ""},
		{"b", 20, 30, ""},
		{"c", 30, 40, ""},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			result, err := doSth(d.in)
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("expected err %s, got err %s", d.errMsg, errMsg)
			}
			if result != d.out {
				t.Errorf("expected %d, got %d", d.out, result)
			}
		})
	}
}
