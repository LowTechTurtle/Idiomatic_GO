package convertbinary

import (
	"encoding/binary"
	"math/bits"
	"unsafe"
)

var isLE bool

func init() {
	var x uint16 = 0xFF00
	xb := *(*[2]byte)(unsafe.Pointer(&x))
	isLE = (xb[0] == 0x00)
}

type Data struct {
	Value  uint32   // 4 bytes
	Label  [10]byte // 10 bytes
	Active bool     //1 byte
	// golang pad 1 more byte
}

const dataSize = unsafe.Sizeof(Data{})

func DataFromBytes(b [dataSize]byte) Data {
	d := Data{}
	d.Value = binary.BigEndian.Uint32(b[:4])
	copy(d.Label[:], b[4:14])
	d.Active = b[14] != 0
	return d
}

func DataFromBytesUnsafe(b [dataSize]byte) Data {
	data := *(*Data)(unsafe.Pointer(&b))
	if isLE {
		data.Value = bits.ReverseBytes32(data.Value)
	}
	return data
}

func BytesFromData(d Data) [dataSize]byte {
	out := [dataSize]byte{}
	binary.BigEndian.PutUint32(out[:4], d.Value)
	copy(out[4:14], d.Label[:])
	if d.Active {
		out[14] = 1
	}
	return out
}

func BytesFromDataUnsafe(d Data) [dataSize]byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	return *(*[dataSize]byte)(unsafe.Pointer(&d))
}

func BytesFromDataUnsafeSlice(d Data) []byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}

	bs := unsafe.Slice((*byte)(unsafe.Pointer(&d)), unsafe.Sizeof(d))
	return bs
}

func DataFromBytesUnsafeSlice(b []byte) Data {
	d := *(*Data)((unsafe.Pointer)(unsafe.SliceData(b)))
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	return d
}
