package chapter_06

import "fmt"

func deferTracing1() {
	b()
}

func trace(s string) {
	fmt.Println("Entering: ", s)
}

func untrace(s string) {
	fmt.Println("Leaving: ", s)
}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}
