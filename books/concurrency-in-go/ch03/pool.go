package ch03

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func pool1() {
	var numCreated int
	pool := sync.Pool{
		New: func() any {
			numCreated++
			a := make([]int, 1024)
			return &a
		},
	}
	var wg sync.WaitGroup
	num := 1024 * 1024
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			a := pool.Get().(*[]int)
			defer pool.Put(a)

			// do something
		}()
	}

	wg.Wait()

	fmt.Println(numCreated, " slices created")
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", ":9999")
		if err != nil {
			log.Fatal(err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			conn.Close()
		}
	}()

	return &wg
}
