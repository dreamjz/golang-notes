package slice

import "fmt"

func newSlice() []int {
	arr := [3]int{1, 2, 3}
	slice := arr[0:1]
	return slice
}

func appendTest1() {
	fmt.Println("------------------------ Int64 8 Byte ------------------------")
	var arr []int64
	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Printf("Arr: %v, Len: %d, Cap: %d\n", arr, len(arr), cap(arr))
	fmt.Println("------------------------ Int32 4 Byte ------------------------")
	var arr2 []int32
	arr2 = append(arr2, 1, 2, 3, 4, 5)
	fmt.Printf("Arr: %v, Len: %d, Cap: %d\n", arr2, len(arr2), cap(arr2))
	fmt.Println("------------------------ Int16 2 Byte ------------------------")
	var arr3 []int16
	arr3 = append(arr3, 1, 2, 3, 4, 5)
	fmt.Printf("Arr: %v, Len: %d, Cap: %d\n", arr3, len(arr3), cap(arr3))
	fmt.Println("------------------------ Int8 1 Byte ------------------------")
	var arr4 []int8
	arr4 = append(arr4, 1, 2, 3, 4, 5)
	fmt.Printf("Arr: %v, Len: %d, Cap: %d\n", arr4, len(arr4), cap(arr4))
}
