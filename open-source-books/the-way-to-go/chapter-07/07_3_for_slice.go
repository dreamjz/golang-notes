package chapter_07

import "fmt"

func forSlice() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	for _, v := range s {
		v++
	}
	fmt.Println(s)
	for i := range s {
		s[i] = 5
	}
	fmt.Println(s)

}
