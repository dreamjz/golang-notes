package natural_sort

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

type generator struct {
	src *rand.Rand
}

func (g *generator) NexInt(max int) int {
	return g.src.Intn(max)
}

func (g *generator) NextString() (str string) {
	strlen := g.src.Intn(6) + 3
	numlen := g.src.Intn(3) + 1
	numpos := g.src.Intn(strlen + 1)
	var num string
	for i := 0; i < numlen; i++ {
		num += strconv.Itoa(g.src.Intn(10))
	}
	for i := 0; i < strlen+1; i++ {
		if i == numpos {
			str += num
		} else {
			str += string('a' + rune(g.src.Intn(16)))
		}
	}
	return str
}

func testSet(seed int) [][]string {
	gen := &generator{
		src: rand.New(rand.NewSource(int64(seed))),
	}
	n := 1000
	if testing.Short() {
		n = 1
	}
	set := make([][]string, n)
	for i := range set {
		strings := make([]string, 10000)
		for idx := range strings {
			strings[idx] = gen.NextString()
		}
		set[i] = strings
	}
	return set
}

func BenchmarkStdStringSort00(b *testing.B) {
	set := testSet(300)
	arr := make([]string, len(set[0]))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, list := range set {
			b.StopTimer()
			copy(arr, list)
			b.StartTimer()
			sort.Strings(arr)
		}
	}
}

func BenchmarkNaturalStringSort01(b *testing.B) {
	set := testSet(300)
	arr := make([]string, len(set[0]))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, list := range set {
			// Resetting the test set to be unsorted does not count.
			b.StopTimer()
			copy(arr, list)
			b.StartTimer()

			sort.Slice(arr, func(i, j int) bool {
				return NaturalLess01(arr[i], arr[j])
			})
		}
	}
}

func BenchmarkNaturalStringSort02(b *testing.B) {
	set := testSet(300)
	arr := make([]string, len(set[0]))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, list := range set {
			// Resetting the test set to be unsorted does not count.
			b.StopTimer()
			copy(arr, list)
			b.StartTimer()

			sort.Slice(arr, func(i, j int) bool {
				return NaturalLess02(arr[i], arr[j])
			})
		}
	}
}
func BenchmarkNaturalStringSort03(b *testing.B) {
	set := testSet(300)
	arr := make([]string, len(set[0]))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, list := range set {
			// Resetting the test set to be unsorted does not count.
			b.StopTimer()
			copy(arr, list)
			b.StartTimer()

			sort.Slice(arr, func(i, j int) bool {
				return NaturalLess03(arr[i], arr[j])
			})
		}
	}
}
