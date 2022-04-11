package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()
	arr := []int{1, 2}
	fmt.Println(arr[3])
}
