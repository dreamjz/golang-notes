package main

import "fmt"

const N = 10000

func main() {
	a := make([]int, N)

	for i := 2; i < N; i++ {
		a[i] = 1
	}

	for i := 2; i < N; i++ {
		for j := i; i*j < N; j++ {
			a[i*j] = 0
		}
	}

	fmt.Print("Result: [")
	for i := range a {
		if a[i] == 1 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("]")
}
