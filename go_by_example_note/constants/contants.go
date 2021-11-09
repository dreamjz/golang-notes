package main

import "fmt"

// const declares a constant value
const s string = "constant"
func main(){
		fmt.Println(s)
		// A const statement can appear anywhere a var statement can
		const n = 50000
		// Constant expression perform arithmetic with arbitrary precision
		const d = 3e10 / n
		fmt.Println(d)
		fmt.Printf("d:%T,%v\n",d,d)
		// A numeric constant has no type until it's given one, such as by
		// explicit conversion
		fmt.Println(int64(d))
}
