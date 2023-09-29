package ch03

import (
	"fmt"
	"sync"
)

func once() {
	var once1, once2 sync.Once
	var count1, count2 int

	incre1 := func() {
		count1++
	}

	incre2 := func() {
		count2++
	}

	decre := func() {
		count2--
	}

	for i := 0; i < 5; i++ {
		once1.Do(incre1)
	}
	fmt.Println(count1)

	once2.Do(incre2)
	once2.Do(decre)

	fmt.Println(count2)
}

func onceDeadlock() {
	var once1, once2 sync.Once

	var f1, f2 func()

	f1 = func() {
		once1.Do(f2)
	}

	f2 = func() {
		once2.Do(f1)
	}

	once1.Do(f1)
	once2.Do(f2)
}
