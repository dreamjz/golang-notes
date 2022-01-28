package main

import (
	"fmt"
	"strings"
	"unsafe"
)

const (
	IntBitSize = int(unsafe.Sizeof(int(0)) * 8)
	N          = 10
	Len        = (N / IntBitSize) + 1
)

type BitMap []int

func (b BitMap) SetBit(k int) {
	i, position := calculateIndexAndOffset(k)
	var flag uint = 1
	b[i] = b[i] | int(flag<<position)
}

func (b BitMap) ClearBit(k int) {
	i, position := calculateIndexAndOffset(k)
	var flag uint = 1
	b[i] = b[i] & int(^(flag << position))
}

func (b BitMap) TestBit(k int) bool {
	i, position := calculateIndexAndOffset(k)
	var flag uint = 1
	return (b[i] & int(flag<<position)) != 0
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

func calculateIndexAndOffset(k int) (int, int) {
	index := k / IntBitSize
	offset := k % IntBitSize
	return index, offset
}

func NewBitMap() BitMap {
	return make(BitMap, Len)
}

func main() {
	fmt.Printf("Int bit size: %d bits\n", IntBitSize)
	bitMap := NewBitMap()
	for i := 2; i < N; i++ {
		bitMap.SetBit(i)
	}
	fmt.Printf("Before: %s", bitMap)
	for i := 2; i < N; i++ {
		for j := i; i*j < N; j++ {
			bitMap.ClearBit(i * j)
		}
	}
	fmt.Print("[")
	for i := 2; i < N; i++ {
		if bitMap.TestBit(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("]")

	fmt.Printf("After: %s\n", bitMap)
	fmt.Printf("Mem usage: [slice descriptor: %d bytes, total: %d bytes]", unsafe.Sizeof(bitMap), unsafe.Sizeof(bitMap)+unsafe.Sizeof([Len]int{}))
}
