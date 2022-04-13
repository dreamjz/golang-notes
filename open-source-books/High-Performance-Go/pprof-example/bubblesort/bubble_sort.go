package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func generateRandIntSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Intn(10000))
	}
	return nums
}

func bubbleSort(nums []int) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func main() {
	f, err := os.OpenFile("./cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln("Create profile error:", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalln("Close file error:", err)
		}
	}()

	if err = pprof.StartCPUProfile(f); err != nil {
		log.Fatalln("Start cpu profile error:", err)
	}
	defer pprof.StopCPUProfile()

	n := 10
	for i := 0; i < 5; i++ {
		nums := generateRandIntSlice(n)
		bubbleSort(nums)
		n *= 10
	}
}
