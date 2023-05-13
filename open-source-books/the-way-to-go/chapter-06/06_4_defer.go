package chapter_06

import "fmt"

func deferF() {
	f1()
}

func f1() {
	fmt.Println("f1 top")
	defer f2()
	fmt.Println("f1 bottom")
}

func f2() {
	fmt.Println("f2")
}
