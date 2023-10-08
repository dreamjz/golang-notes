package main

import "fmt"

var n1, n2 = 10, 20

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dc() {
	if max(n1, n2) == n1 {
		fmt.Println(n1)
	}
}
