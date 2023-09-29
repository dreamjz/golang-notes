package ch04

import (
	"fmt"
	"net/http"
)

func eh1() {
	checkStatus := func(urls []string, done <-chan bool) <-chan *http.Response {
		respch := make(chan *http.Response)
		go func() {
			defer close(respch)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}

				select {
				case <-done:
					return
				case respch <- resp:
				}
			}
		}()

		return respch
	}

	done := make(chan bool)
	defer close(done)

	urls := []string{"https://google.com", "https://badhost"}
	for resp := range checkStatus(urls, done) {
		fmt.Println("Response: ", resp.Status)
	}
}

type Result struct {
	err  error
	resp *http.Response
}

func eh2() {
	checkStatus := func(urls []string, done <-chan bool) <-chan *Result {
		resCh := make(chan *Result)
		go func() {
			defer close(resCh)
			for _, url := range urls {
				resp, err := http.Get(url)

				select {
				case <-done:
					return
				case resCh <- &Result{err, resp}:
				}
			}
		}()

		return resCh
	}

	done := make(chan bool)
	defer close(done)

	urls := []string{"https://google.com", "https://badhost"}
	for res := range checkStatus(urls, done) {
		if res.err != nil {
			fmt.Println("Err: ", res.err)
		} else {
			fmt.Println("Status: ", res.resp.Status)
		}
	}

}
