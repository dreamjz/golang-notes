package slice_operation

func copySlice(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func copySlice2(a []int) []int {
	b := make([]int, len(a))
	b = append([]int(nil), a...)
	return b
}

func copySlice3(a []int) []int {
	b := make([]int, len(a))
	b = append(a[:0:0], a...)
	return b
}
