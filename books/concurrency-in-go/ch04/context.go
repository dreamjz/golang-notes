package ch04

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	LocaleCanceled    = errors.New("locale canceled")
	LocaleUnsupported = errors.New("locale unsupported")
)

func doneChannelPattern() {
	done := make(chan struct{})
	defer close(done)

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		err := printGreeting(done)
		if err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		err := printFarewell(done)
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}

func printFarewell(done <-chan struct{}) error {
	farewell, err := genFarewell(done)
	if err != nil {
		return err
	}
	fmt.Println(farewell, " world")
	return nil
}

func printGreeting(done <-chan struct{}) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return err
	}
	fmt.Println(greeting, " world")
	return nil
}

func genFarewell(done <-chan struct{}) (string, error) {
	switch loc, err := locale(done); {
	case err != nil:
		return "", err
	case loc == "en/us":
		return "goodbye", nil
	default:
		return "", LocaleUnsupported
	}
}

func genGreeting(done <-chan struct{}) (string, error) {
	switch loc, err := locale(done); {
	case err != nil:
		return "", err
	case loc == "en/us":
		return "hello", nil
	default:
		return "", LocaleUnsupported
	}
}

func locale(done <-chan struct{}) (string, error) {
	select {
	case <-done:
		return "", LocaleCanceled
	case <-time.After(3 * time.Second):
	}
	return "en/us", nil
}

func usingCtx() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := printGreetingWithCtx(ctx); err != nil {
			fmt.Println(err)
			cancel() // 取消所有的 goroutine 工作
		}
	}()

	go func() {
		defer wg.Done()
		if err := printFarewellWithCtx(ctx); err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}

func printGreetingWithCtx(ctx context.Context) error {
	greeting, err := genGreetingWithCtx(ctx)
	if err != nil {
		return err
	}
	fmt.Println(greeting, " world")
	return nil
}

func printFarewellWithCtx(ctx context.Context) error {
	farewell, err := genFarewellWithCtx(ctx)
	if err != nil {
		return err
	}
	fmt.Println(farewell, " world")
	return nil
}

func genFarewellWithCtx(ctx context.Context) (string, error) {
	switch loc, err := localeWithCtx(ctx); {
	case err != nil:
		return "", err
	case loc != "en/us":
		return "", LocaleUnsupported
	default:
		return "goodbye", nil
	}
}

func genGreetingWithCtx(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	switch loc, err := localeWithCtx(ctx); {
	case err != nil:
		return "", err
	case loc != "en/us":
		return "", LocaleUnsupported
	default:
		return "hello", nil
	}
}

func localeWithCtx(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(3 * time.Second):
	}
	return "en/us", nil
}
