package _2_datastructure

import (
	"fmt"
	"unsafe"
)

type demo1 struct {
	a int8
	b int16
	c int32
}

type demo2 struct {
	a int8
	b int32
	c int16
}

func memAlign() {
	fmt.Println(unsafe.Sizeof(demo1{}))
	fmt.Println(unsafe.Sizeof(demo2{}))
}
