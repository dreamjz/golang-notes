package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	sliceOfInt := []int{1, 2, 7}
	arrayOfInt := [3]int{8, 12}

	// ToIntSlice
	fmt.Println(cast.ToIntSlice(sliceOfInt)) // [1,3,7]
	fmt.Println(cast.ToIntSlice(arrayOfInt)) // [8,12,0]

	sliceOfInterface := []interface{}{1, 2.0, "kesa"}
	sliceOfString := []string{"a", "b", "cd"}
	stringFields := " abc def hij"
	any := interface{}(37)
	// ToStringSlice
	fmt.Println(cast.ToStringSlice(sliceOfInterface)) // [1 2 kesa]
	fmt.Println(cast.ToStringSlice(sliceOfString))    // [a b cd]
	fmt.Println(cast.ToStringSlice(stringFields))     // [abc def hij]
	fmt.Println(cast.ToStringSlice(any))              // 37

}
