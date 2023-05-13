package chapter_14

import (
	"fmt"
	"time"
)

func mainTimer() {
	tick := time.Tick(time.Second)
	timer := time.After(5 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-timer:
			fmt.Println("boom")
			return
		default:
			fmt.Println(".")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
