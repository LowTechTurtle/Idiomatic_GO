package settear

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("setup for testing")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("tearing down")
	os.Exit(exitVal)
}

func Test_one(t *testing.T) {
	fmt.Println("Test_one use time: ", testTime)
}

func Test_two(t *testing.T) {
	fmt.Println("Test_two use time: ", testTime)
}