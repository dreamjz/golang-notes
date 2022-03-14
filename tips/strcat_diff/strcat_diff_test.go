package strcat_diff

import "testing"

var (
	strsLong   []string
	strsShort  []string
	strsSingle []string
	strsData   []string
)

func init() {
	str := "Golang 字符串拼接!!"
	for i := 0; i < 200; i++ {
		strsLong = append(strsLong, str)
	}
	for i := 0; i < 2; i++ {
		strsShort = append(strsShort, str)
	}
	strsSingle = []string{"Golang 字符串拼接!!"}
	strsData = strsLong
}

func BenchmarkUsingAddOperator(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UsingAddOperator(strsData)
	}
	b.StopTimer()
}

func BenchmarkUsingFmtSprint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		UsingFmtSprint(strsData)
	}
	b.StopTimer()
}

func BenchmarkUsingStringsJoin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UsingStringsJoin(strsData)
	}
	b.StopTimer()
}

func BenchmarkUsingBytesBuffer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UsingBytesBuffer(strsData)
	}
	b.StopTimer()
}

func BenchmarkUsingStringsBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UsingStringsBuilder(strsData)
	}
	b.StopTimer()
}

func BenchmarkUsingStringBuilder2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UsingStringsBuilder2(strsData)
	}
	b.StopTimer()
}
