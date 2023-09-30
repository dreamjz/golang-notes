package ch04

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func findPrime() {
	randIntn := func() int {
		return rand.Intn(5000000000)
	}

	done := make(chan bool)
	defer close(done)
	in := repeatFn(done, randIntn)

	start := time.Now()
	for prime := range take(done, primeFinder(done, in), 10) {
		fmt.Println(prime)
	}
	fmt.Println("Processing time: ", time.Since(start))
}

func findPrimesUsingFaninout() {
	randIntn := func() int {
		return rand.Intn(5000000000)
	}

	done := make(chan bool)
	defer close(done)

	in := repeatFn(done, randIntn)

	// fan out
	numFinders := runtime.NumCPU()
	finders := make([]<-chan int, numFinders)
	for i := range finders {
		finders[i] = primeFinder(done, in)
	}

	start := time.Now()

	// fan in
	for v := range take(done, fanIn(done, finders...), 10) {
		fmt.Println(v)
	}

	fmt.Println("Processing time: ", time.Since(start))
}

func fanIn(done <-chan bool, chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	multiplex := func(c <-chan int) {
		defer wg.Done()
		for v := range c {
			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}

	// get value from every channel
	// and put into out
	wg.Add(len(chs))
	for _, c := range chs {
		go multiplex(c)
	}

	// waiting for reading operations
	// done, then close out channel
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func primeFinder(done <-chan bool, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-done:
				return
			default:
			}

			if isPrime(v) {
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

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	r := int(math.Sqrt(float64(n)))
	for i := 2; i <= r; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
