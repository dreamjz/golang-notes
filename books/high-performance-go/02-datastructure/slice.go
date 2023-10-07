package _2_datastructure

func getLastElemsUsingReslice(a []int, n int) []int {
	return a[len(a)-n:]
}

func getLastElemsUsingCopy(a []int, n int) []int {
	res := make([]int, n)
	copy(res, a[len(a)-n:])
	return res
}
