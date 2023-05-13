package chapter_06

import "fmt"

func printFib() {
	res := 0
	for i := 0; i <= 5; i++ {
		res = fibonacci(i)
		fmt.Printf("%d ", res)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
