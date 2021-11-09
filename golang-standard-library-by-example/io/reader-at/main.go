package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello World")
	p := make([]byte, 10)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s,%d", p, n)
}
