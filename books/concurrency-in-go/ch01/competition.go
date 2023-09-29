package ch01

import "fmt"

func competition() {
	var data int
	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
