package main

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	fmt.Printf("%q\n", randomString(10))
}

func BenchmarkConcatString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatString(1000)
	}
}
