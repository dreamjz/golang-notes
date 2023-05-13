package chpater_10

import "fmt"

type A struct {
	n int
}

func (a *A) set(n int) {
	a.n = n
}

func (a A) String() string {
	return fmt.Sprint(a.n)
}

func methods() {
	// 值
	a1 := A{1}
	fmt.Println(a1)
	a1.set(2)
	fmt.Println(a1)

	// 指针
	a2 := &A{10}
	fmt.Println(a2)
	a2.set(20)
	fmt.Println(a2)
}
