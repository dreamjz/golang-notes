package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("create file failed", err)
	}
	defer file.Close()
	file.WriteString("Hello Alice")
	n, err := file.WriteAt([]byte("World"), 6)
	if err != nil {
		log.Fatal("write failed", err)
	}
	fmt.Println(n)
}
