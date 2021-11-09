package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	// ToString
	fmt.Println(cast.ToString("dreamjz"))          // dreamjz
	fmt.Println(cast.ToString(8))                  // 8
	fmt.Println(cast.ToString(8.31))               // 8.31
	fmt.Println(cast.ToString([]byte("one time"))) // one time
	fmt.Println(cast.ToString(nil))                // ""

	var foo interface{} = "one more time"
	fmt.Println(cast.ToString(foo)) // one more time

	// To int
	fmt.Println(cast.ToInt(8))     // 8
	fmt.Println(cast.ToInt(8.31))  // 8.31
	fmt.Println(cast.ToInt("8"))   // 8
	fmt.Println(cast.ToInt(true))  // 1
	fmt.Println(cast.ToInt(false)) // 0

	var eight interface{} = 8
	fmt.Println(cast.ToInt(eight)) // 8
	fmt.Println(cast.ToInt(nil))   // 0

	// pointer
	p := new(int)
	*p = 8
	fmt.Println(cast.ToInt(p)) // 8

	pp := &p
	fmt.Println(cast.ToInt(pp)) // 8

}
