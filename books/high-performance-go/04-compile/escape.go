package main

import "fmt"

type s struct {
	name string
}

func getS() *s {
	return &s{}
}

func escape() {
	a := getS()
	_ = a
}

func escapeAny() {
	b := 1
	fmt.Println(b)
	d := interface{}(b)
	_ = d
}

func readAny(v interface{}) {
	c := v
	_ = c
}

func localEscape() {
	s1 := make([]int, 0)
	s2 := make([]int, 1024)
	s3 := make([]int, 2048)
	s4 := make([]int, 10000)
	n := 1
	s5 := make([]int, n)

	_ = s1
	_ = s2
	_ = s3
	_ = s4
	_ = s5
}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func closure() {
	incre := Increase()

	incre()
}
