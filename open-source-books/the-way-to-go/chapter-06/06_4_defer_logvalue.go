package chapter_06

import (
	"io"
	"log"
)

func logvalue() {
	f3("GO")
}

func f3(s string) (n int, err error) {
	defer func() {
		log.Printf("f1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
