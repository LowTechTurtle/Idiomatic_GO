package accessunexported

import (
	"fmt"
	"reflect"
	"unsafe"
)

type HasUnexportedField struct {
	A int
	b bool
	C string
}

// normally, the struct should be in another package
// but here i am just demonstrating and im too lazy
// to write it in another package
func SetbUnsafe(huf *HasUnexportedField) {
	sf, _ := reflect.TypeOf(huf).Elem().FieldByName("b")
	offset := sf.Offset
	start := unsafe.Pointer(huf)
	pos := unsafe.Add(start, offset)
	b := (*bool)(pos)
	fmt.Println(b) // read the value
	*b = true // set the value
}
