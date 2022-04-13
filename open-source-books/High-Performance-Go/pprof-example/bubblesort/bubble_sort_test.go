package main

import (
	"fmt"
	"testing"
)

func TestGenerateRandIntSlice(t *testing.T) {
	nums := generateRandIntSlice(10)
	fmt.Println(nums)
}

func TestBubbleSort(t *testing.T) {
	nums := generateRandIntSlice(10)
	fmt.Println("Before:", nums)
	bubbleSort(nums)
	fmt.Println("After:", nums)
}
