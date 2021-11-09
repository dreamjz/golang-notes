package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var err error

func main() {
	fmt.Println("Write to standard output:")
	err = WriteTo(os.Stdout, []byte("string write to stdout"))
	CheckError("write to stdout failed", err)
	fmt.Println()

	fmt.Println("Write to file ... Start")
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	CheckError("crate file failed", err)
	err = WriteTo(file, []byte("string write to file"))
	fmt.Println("Write to file ... Done")
}

func WriteTo(writer io.Writer, data []byte) error {
	_, err := writer.Write(data)
	return err
}

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, ":", err)
	}
}
