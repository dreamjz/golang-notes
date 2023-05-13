package chapter_06

import "fmt"

func varnumpar() {
	x := min(1, 2, 3, 0)
	fmt.Println(x)
	s := []int{5, 3, 4, 1}
	x = min(s...)
	fmt.Println(x)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
