package _2_datastructure

import (
	"fmt"
	"testing"
)

func BenchmarkStrConcat(b *testing.B) {
	tests := []struct {
		name string
		f    func(int, string) string
	}{
		{name: "UsingAddOperator", f: concatStrUsingAdd},
		{name: "UsingSprintf", f: concatStrUsingSprintf},
		{name: "UsingStringsBuilder", f: concatStrUsingStringsBuilder},
		{name: "UsingBuffer", f: concatStrUsingBuffer},
		{name: "UsingByteSlice", f: concatStrUsingByteSlice},
	}

	const (
		StrLen int   = 10
		Seed   int64 = 47
	)

	for k := 0; k <= 100000; k *= 10 {
		for _, tt := range tests {
			b.Run(fmt.Sprintf("%-20s_%.0e", tt.name, float64(k)), func(b *testing.B) {
				str := genStr(StrLen, Seed)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					tt.f(k, str)
				}
			})
		}

		if k == 0 {
			k = 1
		}
	}
}
