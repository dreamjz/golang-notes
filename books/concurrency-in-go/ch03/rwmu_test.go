package ch03

import (
	"fmt"
	"testing"
)

func BenchmarkLock(b *testing.B) {
	tests := []struct {
		name string
		f    func(int)
	}{
		{name: "UsingMutex", f: usingMu},
		{name: "UsingRWMutex", f: usingRWMu},
	}

	for k := 10; k <= 1000; k *= 10 {
		for _, tt := range tests {
			b.Run(fmt.Sprintf("%-15s_%.0e", tt.name, float64(k)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tt.f(k)
				}
			})
		}
	}
}
