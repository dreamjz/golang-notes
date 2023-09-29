package ch03

import (
	"fmt"
	"sync"
	"time"
)

func cond() {
	var wg sync.WaitGroup
	c := sync.NewCond(&sync.Mutex{})

	wg.Add(4)
	done := false
	go reader("reader1", &wg, c, &done)
	go reader("reader2", &wg, c, &done)
	go reader("reader3", &wg, c, &done)

	time.Sleep(3 * time.Second) // 确保消费者已经进入等待状态

	go writer(&wg, c, &done)
	wg.Wait()
}

func reader(name string, wg *sync.WaitGroup, c *sync.Cond, done *bool) {
	defer wg.Done()
	c.L.Lock()
	for !*done {
		fmt.Printf("%s is waiting\n", name)
		c.Wait()
	}
	c.L.Unlock()
	fmt.Printf("%s is reading\n", name)
}

func writer(wg *sync.WaitGroup, c *sync.Cond, done *bool) {
	defer wg.Done()

	c.L.Lock()
	fmt.Println("start writing")
	*done = true
	fmt.Println("done")
	c.L.Unlock()

	fmt.Println("notify all")
	c.Broadcast()

	//通知等待时间最长的那个
	//fmt.Println("notify one")
	//c.Signal()
}
