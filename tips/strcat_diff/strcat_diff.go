package strcat_diff

import (
	"bytes"
	"fmt"
	"strings"
)

// UsingAddOperator use operator '+'
func UsingAddOperator(strs []string) string {
	var s string
	for _, v := range strs {
		s += v
	}
	return s
}

// UsingFmtSprint use fmt.Sprint function
func UsingFmtSprint(strs []string) string {
	var s string
	s = fmt.Sprint(strs)
	return s
}

// UsingStringsJoin use strings.Join
func UsingStringsJoin(strs []string) string {
	var s string
	s = strings.Join(strs, "")
	return s
}

// UsingBytesBuffer use bytes.Buffer
func UsingBytesBuffer(strs []string) string {
	var buf bytes.Buffer
	for _, v := range strs {
		buf.WriteString(v)
	}
	return buf.String()
}

// UsingStringsBuilder use string.Builder
func UsingStringsBuilder(strs []string) string {
	var sb strings.Builder
	for _, v := range strs {
		sb.WriteString(v)
	}
	return sb.String()
}

// UsingStringsBuilder2 use string.Builder
// and calculate the length of final string
func UsingStringsBuilder2(strs []string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}
	var n int
	for _, v := range strs {
		n += len(v)
	}

	var sb strings.Builder
	sb.Grow(n)
	for _, v := range strs {
		sb.WriteString(v)
	}
	return sb.String()
}
