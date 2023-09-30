package ch05

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"os"
	"sync"
)

type APIConnection struct {
	limiter *rate.Limiter
}

func Open() *APIConnection {
	return &APIConnection{
		limiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	if err := a.limiter.WaitN(ctx, 2); err != nil {
		return err
	}

	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.limiter.Wait(ctx); err != nil {
		return err
	}

	return nil
}

func req() {
	defer log.Println("Done.")

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)
	ctx := context.Background()
	var wg sync.WaitGroup

	conn := Open()

	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := conn.ReadFile(ctx)
			if err != nil {
				log.Println(err)
			}
			log.Println("Read file")
		}()

		go func() {
			defer wg.Done()
			err := conn.ResolveAddress(ctx)
			if err != nil {
				log.Println(err)
			}
			log.Println("Resolve address")
		}()
	}

	wg.Wait()
}
