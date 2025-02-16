package main

/*
	extern int mini_calc(char *op, int a, int b);
*/
import "C"
import (
	"fmt"
)

func main() {
	fmt.Println(C.mini_calc(C.CString("+"), 1, 2))
}
