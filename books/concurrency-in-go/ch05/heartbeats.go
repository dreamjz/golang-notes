package ch05

import (
	"fmt"
	"math/rand"
	"time"
)

func heartbeat1() {
	doWork := func(done <-chan struct{}, pulseInterval time.Duration) (<-chan struct{}, <-chan time.Time) {
		heartbeat := make(chan struct{})
		resCh := make(chan time.Time)
		go func() {
			defer close(resCh)
			defer close(heartbeat)

			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				default: // 防止阻塞，当无法发出心跳时
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse:
						sendPulse()
					case resCh <- r:
						return
					}
				}
			}

			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()

		return heartbeat, resCh
	}

	done := make(chan struct{})
	time.AfterFunc(6*time.Second, func() {
		close(done)
	})

	timeout := 2 * time.Second
	heartbeat, resCh := doWork(done, timeout/2)

	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-resCh:
			if !ok {
				return
			}
			fmt.Println(r.Second())
		case <-time.After(timeout):
			return
		}

	}
}

func heartbeat2() {
	doWork := func(done <-chan struct{}, pulseInterval time.Duration) (<-chan struct{}, <-chan time.Time) {
		heartbeat := make(chan struct{})
		resCh := make(chan time.Time)
		go func() {
			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				default: // 防止阻塞，当无法发出心跳时
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse:
						sendPulse()
					case resCh <- r:
						return
					}
				}
			}

			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()

		return heartbeat, resCh
	}

	done := make(chan struct{})
	time.AfterFunc(6*time.Second, func() {
		close(done)
	})

	timeout := 2 * time.Second
	heartbeat, resCh := doWork(done, timeout/2)

	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-resCh:
			if !ok {
				return
			}
			fmt.Println(r.Second())
		case <-time.After(timeout):
			fmt.Println("goroutine is no healthy")
			return
		}

	}
}

func heartbeat3() {
	doWork := func(done <-chan struct{}) (<-chan struct{}, <-chan int) {
		heartbeat := make(chan struct{}, 1) // 缓冲大小为 1
		resCh := make(chan int)
		go func() {
			defer close(resCh)
			defer close(heartbeat)

			for i := 0; i < 5; i++ {
				select {
				case heartbeat <- struct{}{}: // 工作开始前的 心跳
				default:
				}

				select {
				case <-done:
					return
				case resCh <- rand.Intn(20):
				}
			}
		}()

		return heartbeat, resCh
	}

	done := make(chan struct{})
	defer close(done)

	hb, res := doWork(done)
	for {
		select {
		case _, ok := <-hb:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-res:
			if !ok {
				return
			}
			fmt.Println(r)
		}

	}
}

func doWork1(done <-chan struct{}, nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)

		// 模拟运行延迟
		time.Sleep(2 * time.Second)

		for _, v := range nums {
			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}()

	return out
}

func doWork2(done <-chan struct{}, nums []int) (<-chan struct{}, <-chan int) {
	out := make(chan int)
	hb := make(chan struct{}, 1)
	go func() {
		defer close(out)
		defer close(hb)

		// 模拟运行延迟
		time.Sleep(2 * time.Second)

		for _, v := range nums {
			select {
			case hb <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}()

	return hb, out
}

func doWork3(done <-chan struct{}, interval time.Duration, nums []int) (<-chan struct{}, <-chan int) {
	out := make(chan int)
	hb := make(chan struct{})

	go func() {
		defer close(out)
		defer close(hb)

		// 模拟延迟
		time.Sleep(2 * time.Second)

		pulse := time.Tick(interval)

	numLoop:
		for _, v := range nums {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					select {
					case hb <- struct{}{}:
					default: // 心跳信号可跳过
					}
				case out <- v:
					continue numLoop // 继续外层循环
				}
			}
		}
	}()

	return hb, out
}
