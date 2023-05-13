package chapter_04

import (
	"fmt"
	"strings"
)

func indexInString() {
	s := "hello world, hello ..."
	str := "hello"

	fmt.Println(s)
	fmt.Printf("The position of first instance of %s is: %d\n", str, strings.Index(s, str))
	fmt.Printf("The position of last instance of %s is: %d\n", str, strings.LastIndex(s, str))

	s1 := "こんにちは"
	ch := 'ん'

	fmt.Println(s1)
	fmt.Printf("The position of first instance of %c is: %d\n", ch, strings.IndexRune(s1, rune(ch)))

}
