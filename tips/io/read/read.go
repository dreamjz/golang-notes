package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"unsafe"
)

func main() {
	file, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()

	var n16 uint16 = 10
	var n32 uint32 = 10
	var n64 uint64 = 10

	b16 := make([]byte, unsafe.Sizeof(n16))
	b32 := make([]byte, unsafe.Sizeof(n32))
	b64 := make([]byte, unsafe.Sizeof(n64))

	binary.BigEndian.PutUint16(b16, n16)
	binary.BigEndian.PutUint32(b32, n32)
	binary.BigEndian.PutUint64(b64, n64)

	buf := make([]byte, 14)
	copy(buf[0:2], b16)
	copy(buf[2:6], b32)
	copy(buf[6:], b64)
	fmt.Println("Write: ", buf)
	file.Write(buf)

	fmt.Printf("Encode: %v, %v, %v \n", b16, b32, b64)
	d16 := binary.BigEndian.Uint16(b16)
	d32 := binary.BigEndian.Uint32(b32)
	d64 := binary.BigEndian.Uint64(b64)

	fmt.Printf("Decode: %v, %v, %v\n", d16, d32, d64)

	bufR := make([]byte, 14)
	if n, err := file.ReadAt(bufR, 0); err != nil {
		panic(fmt.Sprintf("%d %s", n, err.Error()))
	}

	fmt.Println("Read: ", bufR)

	dn16 := binary.BigEndian.Uint16(bufR[0:2])
	dn32 := binary.BigEndian.Uint32(bufR[2:6])
	dn64 := binary.BigEndian.Uint64(bufR[6:])

	fmt.Printf("File Decode: %v, %v, %v\n", dn16, dn32, dn64)
}
