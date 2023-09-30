package ch05

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func rc() {
	doWork := func(done <-chan struct{}, id int, wg *sync.WaitGroup, out chan<- int) {
		defer wg.Done()

		start := time.Now()
		// 模拟随机延迟
		procTime := time.Duration(rand.Intn(10)+1) * time.Second

		select {
		case <-done:
		case <-time.After(procTime):
		}

		select {
		case <-done:
		case out <- id:
		}

		t := time.Since(start)
		if t < procTime { // 被取消时，预计处理时间
			t = procTime
		}

		fmt.Printf("%d took %v\n", id, t)
	}

	done := make(chan struct{})
	in := make(chan int)
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go doWork(done, i, &wg, in)
	}

	res := <-in
	close(done)
	wg.Wait()

	fmt.Println("Receive result from #", res)
}
