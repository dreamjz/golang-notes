package chapter_13

import (
	"fmt"
	"log"
)

type handlerFunc func(a int, b string)

func recoverClosure() {
	f := errorHandler(handler1)
	f(1, "A")
}

func handler1(a int, b string) {
	panic(fmt.Errorf("handler1 error, a: %d, b: %q", a, b))
}

func errorHandler(fn handlerFunc) handlerFunc {
	return func(a int, b string) {
		defer func() {
			if err, ok := recover().(error); ok {
				log.Printf("run time panic: %v", err)
			}
		}()
		fn(a, b)
	}
}
