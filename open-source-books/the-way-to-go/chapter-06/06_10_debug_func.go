package chapter_06

import (
	"log"
	"runtime"
)

func debugf() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	where()
}
