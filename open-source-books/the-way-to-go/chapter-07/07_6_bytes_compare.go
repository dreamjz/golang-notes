package chapter_07

import "fmt"

func bytcom() {
	a, b := []byte("tell"), []byte("hell")
	fmt.Println(compare(a, b))
}

func compare(a, b []byte) int {
	la, lb := len(a), len(b)
	for i := 0; i < la && i < lb; i++ {
		switch {
		case a[i] < b[i]:
			return -1
		case a[i] > b[i]:
			return 1
		}
	}
	switch {
	case la < lb:
		return -1
	case la > lb:
		return 1
	}
	return 0
}
