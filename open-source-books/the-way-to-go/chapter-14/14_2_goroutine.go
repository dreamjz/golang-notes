package chapter_14

import "fmt"

func gr() {
	ch := make(chan string)
	go sendString(ch)
	go getString(ch)
	
}

func sendString(ch chan string) {
	ch <- "A"
	ch <- "B"
	ch <- "C"
	ch <- "D"
	ch <- "E"
}

func getString(ch chan string) {
	var s string
	for {
		s = <-ch
		fmt.Println(s)
	}
}
