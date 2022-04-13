package generate

import "testing"

const size = 1000000

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(size)
	}
}

func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(size)
	}
}

func benchGenerate(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(i)
	}
}

func BenchmarkGenerateE3(b *testing.B) {
	benchGenerate(1000, b)
}

func BenchmarkGenerateE4(b *testing.B) {
	benchGenerate(10000, b)
}

func BenchmarkGenerateE5(b *testing.B) {
	benchGenerate(100000, b)
}

func BenchmarkGenerateE6(b *testing.B) {
	benchGenerate(1000000, b)
}
