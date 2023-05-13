package chapter_13

import (
	"fmt"
)

func panicRecover() {
	fmt.Println("Calling test")
	test()
	fmt.Println("Test completed")
}

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panicing %s \n", err)
		}
	}()
	badCall()
	fmt.Printf("After bad call \n")
}
