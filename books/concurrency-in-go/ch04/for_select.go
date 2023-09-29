package ch04

import (
	"fmt"
	"math/rand"
	"time"
)

func fs1() {
	doWork := func(strch <-chan string) <-chan bool {
		completed := make(chan bool)
		go func() {
			defer close(completed)
			for str := range strch {
				fmt.Println(str)
			}
		}()

		return completed
	}

	completed := doWork(nil)
	<-completed
}

func fs2() {
	doWork := func(strch <-chan string, done <-chan bool) <-chan bool {
		completed := make(chan bool)
		go func() {
			defer close(completed)
			fmt.Println("goroutine starting ...")
			for {
				select {
				case <-done:
					return
				case str := <-strch:
					fmt.Println(str)
				}
			}
		}()

		return completed
	}

	done := make(chan bool)
	completed := doWork(nil, done)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("Canceling work of goroutine...")
		close(done)
	}()

	<-completed
}

func fs3() {
	randGen := func() <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			defer fmt.Println("Closing output channel...")
			for {
				out <- rand.Int()
			}
		}()

		return out
	}

	in := randGen()
	for i := 0; i < 5; i++ {
		if n, ok := <-in; ok {
			fmt.Println("Receiving: ", n)
		}
	}

	fmt.Println("Receiving done...")
}

func fs4() {
	randGen := func(done <-chan bool) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			defer fmt.Println("Closing output channel...")

			for {
				select {
				case <-done:
					return
				case out <- rand.Int():
				}
			}
		}()

		return out
	}

	done := make(chan bool)
	in := randGen(done)

	for i := 0; i < 5; i++ {
		if n, ok := <-in; ok {
			fmt.Println("Receiving: ", n)
		}
	}

	fmt.Println("Receiving done...")
	close(done)
}
