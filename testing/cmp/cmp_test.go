package persona

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	turtle1 := Person{Name: "turtle",
	Age: 20, DateAdded: time.Now(),}
	turtle2 := CreatePerson("turtle", 20)
	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})
	if diff := cmp.Diff(turtle1, turtle2, comparer); diff != "" {
		t.Error(diff)
	}
}