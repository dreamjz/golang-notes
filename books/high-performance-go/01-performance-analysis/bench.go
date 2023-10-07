package _1_performance_analysis

import (
	"math/rand"
	"time"
)

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func genSlice(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := make([]int, 0)
	for i := 0; i < n; i++ {
		a = append(a, r.Int())
	}
	return a
}

func genSliceWithCap(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, r.Int())
	}
	return a
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
