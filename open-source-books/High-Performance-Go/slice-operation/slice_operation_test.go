package slice_operation

import (
	"math/rand"
	"testing"
	"time"
)

var testSlice []int

const N = 100000000

func init() {
	testSlice = generateSlice(N)
}

func generateSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, n)
	for i := range a {
		a[i] = rand.Intn(n)
	}
	return a
}

func testCopy(f func([]int) []int) []int {
	a := make([]int, len(testSlice))
	copy(a, testSlice)
	return f(a)
}

func intSliceEq(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCopySlice(t *testing.T) {
	r := testCopy(copySlice)
	if !intSliceEq(r, testSlice) {
		t.Error("failed")
	}
}

func TestCopySlice2(t *testing.T) {
	r := testCopy(copySlice2)
	if !intSliceEq(r, testSlice) {
		t.Error("failed")
	}
}

func TestCopySlice3(t *testing.T) {
	r := testCopy(copySlice3)
	if !intSliceEq(r, testSlice) {
		t.Error("failed")
	}
}

func benchCopy(b *testing.B, f func([]int) []int) {
	a := make([]int, len(testSlice))
	copy(a, testSlice)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		f(a)
	}
}

func BenchmarkCopySlice(b *testing.B) {
	benchCopy(b, copySlice)
}

func BenchmarkCopySlice2(b *testing.B) {
	benchCopy(b, copySlice2)
}

func BenchmarkCopySlice3(b *testing.B) {
	benchCopy(b, copySlice3)
}
