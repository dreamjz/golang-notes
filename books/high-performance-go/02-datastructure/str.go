package _2_datastructure

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"unsafe"
)

func getLetter(seed int64) byte {
	r := rand.New(rand.NewSource(seed))
	return byte('a' + r.Intn(26))
}

func genStr(n int, seed int64) string {
	buf := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		buf = append(buf, getLetter(seed))
	}

	return *(*string)(unsafe.Pointer(&buf))
}

func concatStrUsingAdd(n int, str string) string {
	var res string
	for i := 0; i < n; i++ {
		res += str
	}
	return res
}

func concatStrUsingSprintf(n int, str string) string {
	var res string
	for i := 0; i < n; i++ {
		res = fmt.Sprintf("%s%s", res, str)
	}
	return res
}

func concatStrUsingStringsBuilder(n int, str string) string {
	var sb strings.Builder
	sb.Grow(n * len(str)) // 预设置 cap，减少内存分配次数

	for i := 0; i < n; i++ {
		sb.WriteString(str)
	}
	return sb.String()
}

func concatStrUsingBuffer(n int, str string) string {
	var buf bytes.Buffer
	buf.Grow(n * len(str)) // 预设 cap，减少内存分配次数

	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}

	return buf.String()
}

func concatStrUsingByteSlice(n int, str string) string {
	buf := make([]byte, 0, n*len(str)) // 预设 cap，减少内存分配次数
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}

	return *(*string)(unsafe.Pointer(&buf))
}
