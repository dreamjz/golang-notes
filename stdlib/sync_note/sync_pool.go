package main

import (
	"fmt"
	"sync"
)

type MyStruct struct {
	A int
}

func main(){
	fmt.Println("Start ...")

	pool := sync.Pool {
		New: func() interface{}{
			return &MyStruct{
				A: 1,
			}
		},
	}

	testObject := pool.Get().(*MyStruct)
	fmt.Printf("\tTest_Object.A: %d\n",testObject.A)

	pool.Put(testObject)

	fmt.Println("End ...")
}