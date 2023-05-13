package chapter_14

import (
	"fmt"
	"time"
)

func cf() {
	ch := pump2()
	go suck2(ch)
	time.Sleep(time.Second)

}

func pump2() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck2(in chan int) {
	for {
		fmt.Println(<-in)
	}
}
