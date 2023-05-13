package chapter_07

import "fmt"

func strsli() {
	str := "HELLO"
	fmt.Println("str: ", str)
	sli := str[:5]
	// cannot change string
	// sli[0] = 'T'
	fmt.Println(sli)

	// convert to []byte
	sb := []byte(str)
	sb[0] = 'T'
	fmt.Println(string(sb))
}
