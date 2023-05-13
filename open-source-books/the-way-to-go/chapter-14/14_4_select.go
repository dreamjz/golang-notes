package chapter_14

import (
	"fmt"
	"time"
)

func mainSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump3(ch1)
	go pump4(ch2)
	go suck3(ch1, ch2)
	time.Sleep(500 * time.Millisecond)
}

func pump3(out chan<- int) {
	for i := 0; ; i++ {
		out <- i + 5
	}
}

func pump4(out chan<- int) {
	for i := 0; ; i++ {
		out <- i * 2
	}
}

func suck3(in1, in2 <-chan int) {
	for {
		select {
		case u := <-in1:
			fmt.Printf("Receive from ch-1: %d\n", u)
		case v := <-in2:
			fmt.Printf("Receive from ch-2: %d\n", v)
		}
	}
}
