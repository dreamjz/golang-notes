package chapter_14

import (
	"fmt"
	"time"
)

func chblock() {
	ch := make(chan int)
	go pump(ch)
	go suck(ch)
	//fmt.Println(<-ch)
	time.Sleep(time.Second)
	fmt.Println()
}

func pump(out chan int) {
	for i := 0; i < 21; i++ {
		out <- i
	}
}

func suck(in chan int) {
	for {
		fmt.Printf("%d, ", <-in)
	}
}
