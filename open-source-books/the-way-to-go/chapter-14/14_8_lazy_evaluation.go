package chapter_14

import "fmt"

var resume chan int

func mainLazyEvaluation() {
	resume = integers()
	fmt.Println(generate1())
	fmt.Println(generate1())
	fmt.Println(generate1())
}

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generate1() int {
	return <-resume
}
