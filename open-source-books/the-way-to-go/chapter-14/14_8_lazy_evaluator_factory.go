package chapter_14

import "fmt"

type Any interface{}
type EvalFunc func(Any) (Any, Any)

func mainLazyEvaluatorFactory() {
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}
	even := BuildLazyIntEvaluator(evenFunc, 0)

	for i := 0; i < 10; i++ {
		fmt.Printf("%dth even number: %d\n", i, even())
	}
}

func BuidlLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retChan := make(chan Any)
	loopFunc := func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retChan <- retVal
		}
	}
	retFunc := func() Any {
		return <-retChan
	}
	go loopFunc()
	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuidlLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}
