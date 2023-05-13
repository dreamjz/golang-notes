package chapter_06

import (
	"fmt"
	"strings"
)

func factoryFunc() {
	addPng := makeAddSuffix(".png")
	addJpg := makeAddSuffix(".jpg")

	filename := "hero"
	fmt.Println(addPng(filename))
	fmt.Println(addJpg(filename))

}

func makeAddSuffix(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
