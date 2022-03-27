package array

import "fmt"

func accessArrayElement(index int) {
	arr := [3]int{1, 2, 3}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic: %v\n",err)
		}
	}()
	n := arr[index]
	fmt.Printf("arr[%d]: %d", index, n)
}

func outOfRange() int {
	arr := [3]int{1, 2, 3}
	var elem int
	// use variable
	//i := 4
	//elem = arr[i]
	// use literal
	elem = arr[2]
	return elem
}
