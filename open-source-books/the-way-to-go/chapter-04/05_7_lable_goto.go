package chapter_04

import "fmt"

func labelInFor() {

LABEL1:
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if j == 2 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}
