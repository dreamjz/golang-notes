package tips

import (
	"fmt"
	"testing"
)

type Question struct {
	params Param
	ans    int
}

type Param struct {
	nums []int
	val  int
}

var (
	qs []Question
)

func init() {
	qs = []Question{
		{
			params: Param{[]int{1, 2, 3, 4, 5, 7, 5, 8, 5, 6}, 5},
			ans:    7,
		},
	}
}

func TestRemoveElementWithExtraArray(t *testing.T) {
	for _, q := range qs {
		res := removeElementWithExtraArray(q.params.nums, q.params.val)
		if res != q.ans {
			t.Errorf("Wrong Answer: [input]: %v, [output]: %v, [answer]: %v\n", q.params, res, q.ans)
		}
		fmt.Printf("[input]: %v, [output]: %v, [answer]: %v\n", q.params, res, q.ans)
	}
}

func TestRemoveElementDoublePointer(t *testing.T) {
	for _, q := range qs {
		res := removeElementDoublePointer(q.params.nums, q.params.val)
		if res != q.ans {
			t.Errorf("Wrong Answer: [input]: %v, [output]: %v, [answer]: %v\n", q.params, res, q.ans)
		}
		fmt.Printf("[input]: %v, [output]: %v, [answer]: %v\n", q.params, res, q.ans)
	}
}

func BenchmarkRemoveElementWithExtraArray(b *testing.B) {
	nums, val := qs[0].params.nums, qs[0].params.val
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		removeElementWithExtraArray(nums, val)
	}
	b.StopTimer()
}

func BenchmarkRemoveElementDoublePointer(b *testing.B) {
	nums, val := qs[0].params.nums, qs[0].params.val
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		removeElementDoublePointer(nums, val)
	}
	b.StopTimer()
}
