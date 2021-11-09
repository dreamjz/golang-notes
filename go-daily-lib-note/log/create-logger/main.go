package main

import (
	"bytes"
	"fmt"
	"log"
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
	buf := &bytes.Buffer{}
	logger := log.New(buf, "[login]", log.LstdFlags|log.Lmsgprefix)
	logger.Printf("User %s(ID:%d) login ", u.Username, u.ID)
	fmt.Print(buf.String())
}
