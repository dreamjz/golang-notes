package main

import (
	"fmt"
	"unsafe"
)

const N = 10

func main() {
	a := make([]int, N)

	// init with 1
	// assume that all is prime
	for i := 2; i < N; i++ {
		a[i] = 1
	}

	// set to 0 if it is not prime number
	for i := 2; i < N; i++ {
		for j := i; i*j < N; j++ {
			a[i*j] = 0
		}
	}

	fmt.Print("Result: [")
	for i := range a {
		if a[i] == 1 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("]")
	fmt.Printf("Mem usage: [slice descriptor: %d bytes, total: %d bytes]", unsafe.Sizeof(a), unsafe.Sizeof(a)+unsafe.Sizeof([N]int{}))
}
