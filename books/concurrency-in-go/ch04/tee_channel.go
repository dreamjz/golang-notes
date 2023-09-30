package ch04

import (
	"fmt"
	"math/rand"
)

func teeCh() {
	done := make(chan bool)
	defer close(done)

	randGen := func() int {
		return rand.Intn(20)
	}

	out1, out2 := tee(done, take(done, repeatFn(done, randGen), 5))

	for v := range out1 {
		fmt.Printf("out1: %d, out2: %d\n", v, <-out2)
	}
}

func orDoneCh(done <-chan bool, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case <-done:
				case out <- v:
				}
			}
		}
	}()

	return out
}

func tee(done <-chan bool, in <-chan int) (<-chan int, <-chan int) {
	out1, out2 := make(chan int), make(chan int)
	go func() {
		defer close(out1)
		defer close(out2)

		for v := range orDoneCh(done, in) {
			out1A, out2A := out1, out2 // 使用临时变量
			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case out1A <- v:
					out1A = nil // 阻塞写入，以便另一个继续
				case out2A <- v:
					out2A = nil // 阻塞写入，以便另一个继续
				}
			}
		}
	}()

	return out1, out2
}
