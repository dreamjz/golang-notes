package _1_performance_analysis

import "testing"

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}

const GenSliceLen = 500000

func BenchmarkGenSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genSlice(GenSliceLen)
	}
}

func BenchmarkGenSliceWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genSliceWithCap(GenSliceLen)
	}
}

func benchGenSliceHelper(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		genSlice(n)
	}
}

func BenchmarkGenSlice1000(b *testing.B)   { benchGenSliceHelper(1000, b) }
func BenchmarkGenSlice10000(b *testing.B)  { benchGenSliceHelper(10000, b) }
func BenchmarkGenSlice100000(b *testing.B) { benchGenSliceHelper(100000, b) }

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := genSliceWithCap(50000)
		b.StartTimer()
		bubbleSort(nums)
	}
}
