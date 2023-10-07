package _2_datastructure

import (
	"fmt"
	"testing"
)

func BenchmarkIterIntSlice(b *testing.B) {
	tests := []struct {
		name string
		f    func([]int)
	}{
		{name: "UsingFor", f: iterUsingFor},
		{name: "UsingRangeIdx", f: iterUsingRangeIdx},
		{name: "UsingRangeVal", f: iterUsingRangeVal},
	}

	for k := 0; k <= 100000; k *= 10 {
		nums := genSliceWithCap(k)

		for _, tt := range tests {
			b.Run(fmt.Sprintf("%-20s_%.e", tt.name, float64(k)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tt.f(nums)
				}
			})
		}

		if k == 0 {
			k = 1
		}
	}

}

func BenchmarkIterStructSlice(b *testing.B) {
	tests := []struct {
		name string
		f    func([]MySt)
	}{
		{name: "UsingFor", f: iterMsUsingFor},
		{name: "UsingRangeIdx", f: iterMsUsingRangeIdx},
		{name: "UsingRangeVal", f: iterMsUsingRangeVal},
	}

	for k := 0; k <= 1000; k *= 10 {
		a := make([]MySt, k)
		for _, tt := range tests {
			b.Run(fmt.Sprintf("%-20s_%.e", tt.name, float64(k)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tt.f(a)
				}
			})
		}

		if k == 0 {
			k = 1
		}
	}

}

func genMsSlice(n int) []*MySt {
	res := make([]*MySt, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, &MySt{})
	}
	return res
}

func BenchmarkIterStructPtrSlice(b *testing.B) {
	tests := []struct {
		name string
		f    func([]*MySt)
	}{
		{name: "UsingFor", f: iterMsPtrUsingFor},
		{name: "UsingRangeIdx", f: iterMsPtrUsingRangeIdx},
		{name: "UsingRangeVal", f: iterMsPtrUsingRangeVal},
	}

	for k := 0; k <= 1000; k *= 10 {
		a := genMsSlice(k)
		for _, tt := range tests {
			b.Run(fmt.Sprintf("%-20s_%.e", tt.name, float64(k)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tt.f(a)
				}
			})
		}

		if k == 0 {
			k = 1
		}
	}

}
