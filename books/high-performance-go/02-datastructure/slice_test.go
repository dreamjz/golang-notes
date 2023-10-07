package _2_datastructure

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
)

func genSliceWithCap(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, rand.Int())
	}

	return res
}

func printMem(t *testing.T) {
	t.Helper()
	var mstat runtime.MemStats
	runtime.ReadMemStats(&mstat)
	t.Logf("%.2f MB", float64(mstat.Alloc)/1024./1024.)
}

func TestGetSlice(t *testing.T) {
	tests := []struct {
		name string
		f    func([]int, int) []int
	}{
		{name: "UsingReslice", f: getLastElemsUsingReslice},
		{name: "UsingCopy", f: getLastElemsUsingCopy},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%-20s", tt.name), func(t *testing.T) {
			var res [][]int
			for i := 0; i < 100; i++ {
				nums := genSliceWithCap(128 * 1024)
				res = append(res, tt.f(nums, 2))
				runtime.GC()
			}
			printMem(t)
		})
	}
}
