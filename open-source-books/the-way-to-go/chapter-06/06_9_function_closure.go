package chapter_06

import (
	"fmt"
)

func fClosure() {
	f := adder2()
	fmt.Println(f(1))
	fmt.Println(f(10))
	fmt.Println(f(100))
}

func adder2() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
