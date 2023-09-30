package ch04

import (
	"fmt"
	"math/rand"
)

func gen1() {
	done := make(chan bool)
	defer close(done)

	for v := range take(done, repeat(done, 1), 3) {
		fmt.Println(v)
	}
}

func gen2() {
	genRand := func() int {
		return rand.Intn(100)
	}

	done := make(chan bool)
	defer close(done)

	for v := range take(done, repeatFn(done, genRand), 5) {
		fmt.Println(v)
	}
}

func take(done <-chan bool, in <-chan int, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case out <- <-in:
			}
		}
	}()

	return out
}

func repeat(done <-chan bool, val int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- val:
			}
		}
	}()

	return out
}

func repeatFn(done <-chan bool, fn func() int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- fn():
			}
		}
	}()

	return out
}
