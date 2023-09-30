package ch05

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"os"
	"sort"
	"sync"
	"time"
)

type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

type MultiLimiter struct {
	limiters []RateLimiter
}

func NewMultiLimiter(lims ...RateLimiter) *MultiLimiter {
	// 按照 limit 升序排列
	sort.Slice(lims, func(i, j int) bool {
		return lims[i].Limit() < lims[j].Limit()
	})
	return &MultiLimiter{limiters: lims}
}

func (m *MultiLimiter) Wait(ctx context.Context) error {
	for i := range m.limiters {
		if err := m.limiters[i].Wait(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (m *MultiLimiter) Limit() rate.Limit {
	return m.limiters[0].Limit()
}

func mullim() {
	defer log.Println("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LUTC | log.Ltime)

	conn := OpenConn()
	var wg sync.WaitGroup

	ctx := context.Background()
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			if err := conn.ReadFile(ctx); err != nil {
				log.Println(err)
			}
			log.Printf("#%d Read file\n", id)
		}(i)

		go func(id int) {
			defer wg.Done()
			if err := conn.ResolveAddress(ctx); err != nil {
				log.Println(err)
			}
			log.Printf("#%d Resolve address\n", id)
		}(i)
	}

	wg.Wait()
}

type APIConn struct {
	limiter RateLimiter
}

func OpenConn() *APIConn {
	// 2 per sec
	secondLimiter := rate.NewLimiter(Per(2, time.Second), 2)
	// 10 per min
	minuteLimiter := rate.NewLimiter(Per(10, time.Minute), 10)

	return &APIConn{
		limiter: NewMultiLimiter(secondLimiter, minuteLimiter),
	}
}

func (a *APIConn) ReadFile(ctx context.Context) error {
	if err := a.limiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func (a *APIConn) ResolveAddress(ctx context.Context) error {
	if err := a.limiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func Per(n int, d time.Duration) rate.Limit {
	return rate.Every(d / time.Duration(n))
}

type APIConn1 struct {
	diskLimiter,
	netLimiter,
	apiLimiter RateLimiter
}

func OpenAPIConn1() *APIConn1 {
	// api limiter
	// 2 per sec
	secondLimiter := rate.NewLimiter(Per(2, time.Second), 2)
	// 10 per min
	minuteLimiter := rate.NewLimiter(Per(10, time.Minute), 10)

	// disk limiter
	diskLimiter := rate.NewLimiter(rate.Limit(1), 1)
	// net limiter
	netLimiter := rate.NewLimiter(Per(3, time.Second), 3)

	return &APIConn1{
		diskLimiter: diskLimiter,
		netLimiter:  netLimiter,
		apiLimiter:  NewMultiLimiter(secondLimiter, minuteLimiter),
	}
}

func (a *APIConn1) ReadFile1(ctx context.Context) error {
	if err := NewMultiLimiter(a.apiLimiter, a.diskLimiter).Wait(ctx); err != nil {
		return err
	}
	return nil
}

func (a *APIConn1) ResolveAddress1(ctx context.Context) error {
	if err := NewMultiLimiter(a.apiLimiter, a.netLimiter).Wait(ctx); err != nil {
		return err
	}
	return nil
}

func mullim1() {
	defer log.Println("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LUTC | log.Ltime)

	conn := OpenAPIConn1()
	var wg sync.WaitGroup

	ctx := context.Background()
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			if err := conn.ReadFile1(ctx); err != nil {
				log.Println(err)
			}
			log.Printf("#%d Read file\n", id)
		}(i)

		go func(id int) {
			defer wg.Done()
			if err := conn.ResolveAddress1(ctx); err != nil {
				log.Println(err)
			}
			log.Printf("#%d Resolve address\n", id)
		}(i)
	}

	wg.Wait()
}
