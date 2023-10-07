package _2_datastructure

func iterUsingFor(a []int) {
	var res int
	for i := 0; i < len(a); i++ {
		res = a[i]
	}
	_ = res
}

func iterUsingRangeIdx(a []int) {
	var res int
	for i := range a {
		res = a[i]
	}
	_ = res
}

func iterUsingRangeVal(a []int) {
	var res int
	for _, v := range a {
		res = v
	}
	_ = res
}

type MySt struct {
	id   int
	data [1024]int
}

func iterMsUsingFor(a []MySt) {
	var res MySt
	for i := 0; i < len(a); i++ {
		res = a[i]
	}
	_ = res
}

func iterMsUsingRangeIdx(a []MySt) {
	var res MySt
	for i := range a {
		res = a[i]
	}
	_ = res
}

func iterMsUsingRangeVal(a []MySt) {
	var res MySt
	for _, v := range a {
		res = v
	}
	_ = res
}

func iterMsPtrUsingFor(a []*MySt) {
	var res *MySt
	for i := 0; i < len(a); i++ {
		res = a[i]
	}
	_ = res
}

func iterMsPtrUsingRangeIdx(a []*MySt) {
	var res *MySt
	for i := range a {
		res = a[i]
	}
	_ = res
}

func iterMsPtrUsingRangeVal(a []*MySt) {
	var res *MySt
	for _, v := range a {
		res = v
	}
	_ = res
}
