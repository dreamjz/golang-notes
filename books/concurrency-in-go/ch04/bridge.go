package ch04

import "fmt"

func bc() {
	genCh := func() <-chan <-chan int {
		out := make(chan (<-chan int))
		go func() {
			defer close(out)
			for i := 0; i < 5; i++ {
				ch := make(chan int, 1)
				ch <- i
				close(ch)
				out <- ch
			}
		}()
		return out
	}

	done := make(chan bool)
	defer close(done)
	for v := range bridge(done, genCh()) {
		fmt.Print(v, " ")
	}
}

func bridge(done <-chan bool, chIn <-chan <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			// 从通道的通道中获取通道
			var c <-chan int
			select {
			case <-done:
				return
			case vc, ok := <-chIn:
				if !ok {
					return
				}
				c = vc
			}

			// 遍历通道，获取值
			for v := range c {
				select {
				case <-done:
					return
				case out <- v:
				}
			}
		}
	}()

	return out
}
