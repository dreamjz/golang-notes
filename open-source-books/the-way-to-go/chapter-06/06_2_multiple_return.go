package chapter_06

import "fmt"

func mutlipleReturn() {
	num := 10

	x2, x3 := getX2AndX3(num)
	printValues(num, x2, x3)

	x2, x3 = getX2AndX3_2(num)
	printValues(num, x2, x3)

}

func printValues(n, v1, v2 int) {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", n, v1, v2)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	fmt.Printf("init value x2: %d, x3: %d\n", x2, x3)
	x2 = 2 * input
	x3 = 3 * input
	return
}
