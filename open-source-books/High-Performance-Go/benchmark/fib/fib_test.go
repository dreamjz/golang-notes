package fib

import (
	"testing"
	"time"
)

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}

func BenchmarkFib2(b *testing.B) {
	time.Sleep(time.Second * 1)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}

func BenchmarkFib3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		time.Sleep(time.Millisecond * 10)
		b.StartTimer()
		fib(30)
	}
}

func BenchmarkFib4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		time.Sleep(time.Millisecond * 10)
		fib(30)
	}
}

func BenchmarkFibWithPreparation(b *testing.B) {
	time.Sleep(time.Second * 1)
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}
