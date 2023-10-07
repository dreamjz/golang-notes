package _2_datastructure

import (
	"fmt"
	"unsafe"
)

func emptyStructMemUsage() {
	es := struct{}{}
	fmt.Println(unsafe.Sizeof(es))
}

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func setOp() {
	s := make(Set)
	s.Add("A")
	s.Add("B")
	s.Add("C")

	fmt.Println(s)
	fmt.Println(s.Has("C"))
	fmt.Println(s.Has("D"))

	s.Delete("B")
	fmt.Println(s)
}
