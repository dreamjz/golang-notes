package ch04

import (
	"fmt"
	"time"
)

func or(chans ...<-chan bool) <-chan bool {
	switch len(chans) { // 递归出口
	case 0:
		return nil
	case 1:
		return chans[0]
	}

	orDone := make(chan bool)
	go func() {
		defer close(orDone)
		select {
		case <-chans[0]:
		case <-chans[1]:
		case <-or(append(chans[2:], orDone)...): // 递归的创建done-chan链
		}
	}()

	return orDone
}

func orDone() {
	sig := func(after time.Duration) <-chan bool {
		done := make(chan bool)
		go func() {
			defer close(done)
			time.Sleep(after)
		}()

		return done
	}

	start := time.Now()
	<-or(
		sig(time.Hour),
		sig(time.Minute),
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(time.Second),
	)

	fmt.Println("Done after: ", time.Since(start))
}
