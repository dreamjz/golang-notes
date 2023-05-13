package chapter_06

import "fmt"

func functionReturn() {
	p2 := add2()
	fmt.Printf("p2(3): %d \n", p2(3))

	twoAdder := adder(2)
	fmt.Printf("twoAdder(3): %d\n", twoAdder(3))
}

func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
