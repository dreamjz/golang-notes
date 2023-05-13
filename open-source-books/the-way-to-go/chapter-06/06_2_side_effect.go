package chapter_06

import "fmt"

func sideEffect() {
	a, b := 2, 3
	var reply int
	multiply(a, b, &reply)
	fmt.Printf("%d * %d = %d\n", a, b, reply)
}

func multiply(a, b int, reply *int) {
	*reply = a * b
}
