package ch05

import (
	"log"
	"os"
	"time"
)

type StartGoroutineFn func(done <-chan struct{}, pulseInterval time.Duration) <-chan struct{}

func NewSteward(timeout time.Duration, startGoroutine StartGoroutineFn) StartGoroutineFn {
	return func(done <-chan struct{}, pulseInterval time.Duration) <-chan struct{} {
		hb := make(chan struct{})
		go func() {
			defer close(hb)

			var wardDone chan struct{}
			var wardHb <-chan struct{}

			startWard := func() { // 启动 goroutine 函数
				wardDone = make(chan struct{})
				wardHb = startGoroutine(or(wardDone, done), timeout/2) // 管理员被关闭时，监控区也关闭
			}

			startWard() // 启动 goroutine

			pulse := time.Tick(pulseInterval)

		monitorLoop:
			for {
				timeoutSignal := time.After(timeout) // 监控超时

				for {
					select {
					case <-pulse: // 管理员心跳
						select {
						case hb <- struct{}{}: // 发送管理员心跳
						default:
						}
					case <-wardHb: // 读取 监控区 心跳
						continue monitorLoop
					case <-timeoutSignal: // 监控超时，重启被监控的 goroutine
						log.Println("steward: ward unhealthy; restarting")
						close(wardDone) // 关闭监控区 goroutine
						startWard()     // 启动
						continue monitorLoop
					case <-done: // 管理员被关闭
						return
					}
				}
			}
		}()
		return hb
	}
}

func or(dones ...<-chan struct{}) <-chan struct{} {
	switch len(dones) {
	case 0:
		return nil
	case 1:
		return dones[0]
	}

	orDone := make(chan struct{})
	go func() {
		defer close(orDone)
		select {
		case <-dones[0]:
		case <-dones[1]:
		case <-or(append(dones[2:], orDone)...):
		}
	}()

	return orDone
}

func monitor() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	// 监控区 goroutine
	// 什么也不做，等待被关闭
	doWork := func(done <-chan struct{}, _ time.Duration) <-chan struct{} {
		log.Println("ward: hello, I'm irresponsible...")
		go func() {
			<-done
			log.Println("ward: I'm halting")
		}()
		return nil
	}

	// 管理员函数，设置 4 秒的超时
	doWorkWithSteward := NewSteward(4*time.Second, doWork)

	// 9 秒后关闭管理员和监控区
	done := make(chan struct{})
	time.AfterFunc(9*time.Second, func() {
		log.Println("Halting steward and ward")
		close(done)
	})

	// 监控管理员心跳
	for range doWorkWithSteward(done, 4*time.Second) {
	}

	log.Println("Done")
}
