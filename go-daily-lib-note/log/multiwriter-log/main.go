package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type User struct {
	ID       int
	Username string
}

func main() {
	u := User{
		ID:       1,
		Username: "Kesa",
	}

	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("./multi.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("create file failed:", err)
	}
	logger := log.New(io.MultiWriter(writer1, writer2, writer3), "[Multi]", log.LstdFlags|log.Lmsgprefix)
	logger.Printf("%s login,ID:%d", u.Username, u.ID)
	fmt.Print("Buf:", writer1.String())
}
