package chapter_04

import (
	"fmt"
	"strings"
)

func presuffix() {
	str := "This is a string"
	prefix := "This"
	suffix := "string"
	
	fmt.Printf("%s has prefix %s? %t \n", str, prefix, strings.HasPrefix(str, prefix))
	fmt.Printf("%s has suffix %s? %t \n", str, suffix, strings.HasSuffix(str, suffix))
}
