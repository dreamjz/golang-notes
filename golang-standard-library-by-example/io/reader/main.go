package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	data []byte
	err  error
)

func main() {
	fmt.Println("Read from Standard input")
	data, err = ReadFrom(os.Stdin, 5)
	CheckError("read from stdin failed", err)
	fmt.Println("  Data from stdin:", string(data))

	fmt.Println("Read from file")
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	defer file.Close()
	CheckError("open file failed", err)
	data, err = ReadFrom(file, 20)
	CheckError("read from file failed", err)
	fmt.Println("  Data from file:", string(data))

	fmt.Println("Read from string")
	data, err = ReadFrom(strings.NewReader("from string"), 20)
	CheckError("read from string failed", err)
	fmt.Println("  Data from string:", string(data))

}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p, nil
	}
	return p, err
}

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, ":", err)
	}
}
