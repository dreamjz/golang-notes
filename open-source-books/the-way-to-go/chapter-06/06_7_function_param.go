package chapter_06

import "fmt"

func funcParam() {
	callBack(1, Add)
}

func Add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func callBack(n int, f func(int, int)) {
	f(n, 2)
}
