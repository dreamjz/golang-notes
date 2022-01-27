package main

import (
	"fmt"
	"strings"
	"unsafe"
)

const (
	IntBitSize = int(unsafe.Sizeof(int(0)) * 8)
	N          = 1000
	Len        = (N / IntBitSize) + 1
)

type BitMap []int

func (b BitMap) SetBit(k int) {
	i := k / IntBitSize
	position := k % IntBitSize

	var flag uint = 1
	flag = flag << position
	b[i] = b[i] | int(flag)
}

func (b BitMap) ClearBit(k int) {
	i := k / IntBitSize
	position := k % IntBitSize

	var flag uint = 1
	flag = ^(flag << position)
	b[i] = b[i] & int(flag)
}

func (b BitMap) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	for i := range b {
		builder.WriteString(fmt.Sprintf("\n\t%s,", SprintBit(b[i])))
	}
	builder.WriteString("\n]")
	return builder.String()
}

func SprintBit(n int) string {
	var builder strings.Builder
	builder.WriteString("[")
	for i := IntBitSize - 1; i >= 0; i-- {
		builder.WriteString(fmt.Sprintf("%d", (n>>i)&1))
	}
	builder.WriteString("]")
	return builder.String()
}

func NewBitMap(cap int) BitMap {
	return make(BitMap, Len)
}

func main() {
	fmt.Printf("Int bit size: %d bits\n", IntBitSize)
	bitMap := NewBitMap(N)
	fmt.Println(SprintBit(7))
	for i := 0; i < N; i++ {
		bitMap.SetBit(i)
	}
	fmt.Println(bitMap)

}
