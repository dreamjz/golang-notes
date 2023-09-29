package ch01

import (
	"fmt"
	"sync"
	"time"
)

func deadLock() {
	type val struct {
		mu sync.Mutex
		v  int
	}

	var wg sync.WaitGroup
	printSum := func(v1, v2 *val) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second) // 尽可能让死锁发生的概率提高，因为无法保证 goroutine 的执行顺序

		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum = %d\n", v1.v+v2.v)
	}

	var a, b val
	wg.Add(2)
	go printSum(&a, &b) // 锁定 a, 等待 b
	go printSum(&b, &a) // 锁定 b, 等待 a
	wg.Wait()
}
