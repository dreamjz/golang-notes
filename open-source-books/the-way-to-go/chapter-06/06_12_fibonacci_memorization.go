package chapter_06

import (
	"fmt"
	"time"
)

const LIMIT = 50

var fibs [LIMIT]uint64

func fibF() {
	timeSpend(fib, printFibs)
	timeSpend(fibMem, printFibs)
}

func fib(n int) (res uint64) {
	if n <= 1 {
		res = 1
		return
	}
	res = fib(n-1) + fib(n-2)
	return
}

func fibMem(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibMem(n-1) + fibMem(n-2)
	}
	// save result
	fibs[n] = res
	return
}

func printFibs(fib func(int) uint64) {
	for i := 0; i < LIMIT; i++ {
		fmt.Printf("%d ", fib(i))
	}
	fmt.Println()
}

func timeSpend(fib func(int) uint64, p func(fib func(int) uint64)) {
	start := time.Now()
	p(fib)
	end := time.Now()
	fmt.Println("Time: ", end.Sub(start))
}
