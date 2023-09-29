package ch04

import "fmt"

func batchProcess() {
	multiply := func(vals []int, multiplier int) []int {
		muls := make([]int, len(vals))
		for i, v := range vals {
			muls[i] = v * multiplier
		}
		return muls
	}

	add := func(vals []int, additive int) []int {
		adds := make([]int, len(vals))
		for i, v := range vals {
			adds[i] = v + additive
		}
		return adds
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Printf("%d, ", v)
	}
}

func streamProcess() {
	
}
