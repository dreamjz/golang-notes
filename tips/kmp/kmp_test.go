package kmp

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unsafe"
)

const (
	N = 100000
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	src = rand.NewSource(time.Now().UnixNano())
	s   string
	p   string
)

func init() {
	s = RandStringBytesMaskImprSrcUnsafe(N)
	p = RandStringBytesMaskImprSrcUnsafe(rand.Intn(N))
}

func TestBruteForce(t *testing.T) {
	res := bruteForce(s, p)
	fmt.Println(res)
}

func TestKmp(t *testing.T) {
	res := kmp(s, p)
	fmt.Println(res)
}

func BenchmarkBruteForce(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bruteForce(s, p)
	}
	b.StopTimer()
}

func BenchmarkKmp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		kmp(s, p)
	}
	b.StopTimer()
}

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
