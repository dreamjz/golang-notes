package _3_concurrency

import (
	"sync"
	"time"
)

type RW interface {
	Write()
	Read()
}

const cost = time.Microsecond

type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	_ = l.count
	time.Sleep(cost)
	l.mu.Unlock()
}

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (rw *RWLock) Write() {
	rw.mu.Lock()
	rw.count++
	time.Sleep(cost)
	rw.mu.Unlock()
}

func (rw *RWLock) Read() {
	rw.mu.RLock()
	_ = rw.count
	time.Sleep(cost)
	rw.mu.RUnlock()
}
