package ch04

import "fmt"

func batchProcess() {
	multiply := func(vals []int, multiplier int) []int {
		muls := make([]int, len(vals))
		for i, v := range vals {
			muls[i] = v * multiplier
		}
		return muls
	}

	add := func(vals []int, additive int) []int {
		adds := make([]int, len(vals))
		for i, v := range vals {
			adds[i] = v + additive
		}
		return adds
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Printf("%d, ", v)
	}
}

func streamProcess() {
	multiply := func(val int, mul int) int {
		return val * mul
	}

	add := func(val int, additive int) int {
		return val + additive
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Print(add(multiply(v, 2), 1), ", ")
	}
}

func streamProcWithChan() {
	intGen := func(done <-chan bool, vals ...int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, v := range vals {
				select {
				case <-done:
					return
				case out <- v:
				}
			}
		}()

		return out
	}

	multiply := func(done <-chan bool, in <-chan int, mul int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for v := range in {
				select {
				case <-done:
					return
				case out <- v * mul:
				}
			}
		}()

		return out
	}

	add := func(done <-chan bool, in <-chan int, additive int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for v := range in {
				select {
				case <-done:
					return
				case out <- v + additive:
				}
			}
		}()

		return out
	}

	done := make(chan bool)
	defer close(done)

	for v := range add(done, multiply(done, intGen(done, 1, 2, 3, 4), 2), 1) {
		fmt.Println(v)
	}
}
