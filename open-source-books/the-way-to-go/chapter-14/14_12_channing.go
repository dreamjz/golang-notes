package chapter_14

import "fmt"

const N = 100000

func mainChanning() {
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < N; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0
	x := <-leftmost
	fmt.Println(x)
}

func f(left, right chan int) {
	left <- 1 + <-right
}
