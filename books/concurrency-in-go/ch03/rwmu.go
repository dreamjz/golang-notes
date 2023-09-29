package ch03

import (
	"sync"
	"time"
)

func usingRWMu(count int) {
	var wg sync.WaitGroup
	wg.Add(count + 1)

	var rw sync.RWMutex
	go producer(&wg, &rw, count)
	for i := 0; i < count; i++ {
		go consumer(&wg, rw.RLocker())
	}

	wg.Wait()
}

func usingMu(count int) {
	var wg sync.WaitGroup
	wg.Add(count + 1)

	var mu sync.Mutex
	go producer(&wg, &mu, count)
	for i := 0; i < count; i++ {
		go consumer(&wg, &mu)
	}

	wg.Wait()
}

func producer(wg *sync.WaitGroup, lock sync.Locker, count int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		lock.Lock()
		lock.Unlock()
	}
}

func consumer(wg *sync.WaitGroup, lock sync.Locker) {
	defer wg.Done()
	defer lock.Unlock()
	lock.Lock()
	time.Sleep(1)
}
