package chapter_14

import "fmt"

func dl() {
	ch := make(chan int)
	ch <- 1
	go fn(ch)
}

func fn(in chan int) {
	fmt.Println(<-in)
}
