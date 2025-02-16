package main

import (
	"fmt"
	"unsafe"
)

type OrderInfo struct {
	OrderCode   rune
	Amount      int
	OrderNumber uint16
	Items       []string
	IsReady     bool
}

type OrderInfoSmol struct {
	Items       []string // 24 bits
	Amount      int      // should 64 bit, 8 bytes on modern system but it could be 32 bits
	OrderCode   rune     // alias for int32 => 4 bytes
	OrderNumber uint16   // 2 bytes
	IsReady     bool     // 1 bit is enough for bool but its size is 1 byte
	// all of this use 24 + 8 + 7 bytes + 1 padded bytes = 40 bytes
}

func main() {
	fmt.Println("OrderInfo size and offset")
	oi := OrderInfo{}
	fmt.Println(unsafe.Sizeof(oi))
	fmt.Println(unsafe.Offsetof(oi.OrderCode))
	fmt.Println(unsafe.Offsetof(oi.Amount))
	fmt.Println(unsafe.Offsetof(oi.OrderNumber))
	fmt.Println(unsafe.Offsetof(oi.Items))
	fmt.Println(unsafe.Offsetof(oi.IsReady))
	fmt.Println("\nOrderInfoSmol size and offset")
	oim := OrderInfoSmol{}
	fmt.Println(unsafe.Sizeof(oim))
	fmt.Println(unsafe.Offsetof(oim.OrderCode))
	fmt.Println(unsafe.Offsetof(oim.Amount))
	fmt.Println(unsafe.Offsetof(oim.OrderNumber))
	fmt.Println(unsafe.Offsetof(oim.Items))
	fmt.Println(unsafe.Offsetof(oim.IsReady))
}
