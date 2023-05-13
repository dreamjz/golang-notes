package chapter_06

import "fmt"

func deferTracing2() {
	b2()
}

func trace2(s string) string {
	fmt.Println("Entering: ", s)
	return s
}

func untrace2(s string) {
	fmt.Println("Leaving: ", s)
}

func a2() {
	defer untrace2(trace2("a"))
	fmt.Println("in a")
}

func b2() {
	defer untrace2(trace2("b"))
	fmt.Println("in b")
	a2()
}
