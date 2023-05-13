package chapter_14

func channelSync(n int) {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}

func channelBuffered(n int) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}
