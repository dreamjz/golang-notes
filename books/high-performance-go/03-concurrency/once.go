package _3_concurrency

import (
	"log"
	"sync"
)

var readCfgOnce sync.Once

func ReadConfig() {
	readCfgOnce.Do(func() {
		log.Println("Reading Config...")
	})
}

func mainFunc() {
	var wg sync.WaitGroup
	n := 10

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			ReadConfig()
		}()
	}

	wg.Wait()
}
