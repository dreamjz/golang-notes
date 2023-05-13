package chapter_14

import "fmt"

func mainFibEvaluator() {
	fibFunc := func(state Any) (Any, Any) {
		os := state.([]uint64)
		v1 := os[0]
		v2 := os[1]
		ns := []uint64{v2, v1 + v2}
		return v1, ns
	}
	fib := BuildLazyUint64Evaluator(fibFunc, []uint64{0, 1})

	for i := 0; i < 10; i++ {
		fmt.Printf("%dth fib : %d\n", i, fib())
	}
}

func BuildLazyUint64Evaluator(evalFunc EvalFunc, initState Any) func() uint64 {
	ef := BuidlLazyEvaluator(evalFunc, initState)
	return func() uint64 {
		return ef().(uint64)
	}
}
