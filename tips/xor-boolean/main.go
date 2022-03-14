package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	// bitwise XOR
	// 0000 ... 0101
	i := 5
	// 0000 ... 0001
	j := 3
	r := i ^ j
	fmt.Printf("%3d: %s\n", i, SprintBits(i))
	fmt.Printf("%3d: %s\n", j, SprintBits(j))
	fmt.Printf("%3d: %s\n", r, SprintBits(r))
	// boolean xor
	fmt.Printf("%5v xor %5v: %v\n", true, true, xor(true, true))
	fmt.Printf("%5v xor %5v: %v\n", true, false, xor(true, false))
	fmt.Printf("%5v xor %5v: %v\n", false, true, xor(false, true))
	fmt.Printf("%5v xor %5v: %v\n", false, false, xor(false, false))
}

func SprintBits(n int) string {
	var sb strings.Builder
	intBitSize := int(unsafe.Sizeof(int(0)) * 8)
	sb.WriteString("[")
	for i := intBitSize - 1; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%d", (n>>i)&1))
	}
	sb.WriteString("]")
	return sb.String()
}

func xor(x, y bool) bool {
	return x != y
}
