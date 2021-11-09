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
	user := User{
		ID:       1,
		Username: "Kesa",
	}

	// set prefix
	log.SetPrefix("[Login]")
	// set flag
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmsgprefix)
	// get flags
	fmt.Println(LogFlagsToString(log.Flags()))
	log.Printf("User: %s login,ID:%d", user.Username, user.ID)
}

func LogFlagsToString(flags int) string {
	var buffer bytes.Buffer

	if flags&log.Ldate == log.Ldate {
		buffer.WriteString("|Ldate")
	}
	if flags&log.Ltime == log.Ltime {
		buffer.WriteString("|Ltime")
	}
	if flags&log.Lmicroseconds == log.Lmicroseconds {
		buffer.WriteString("|Lmicroseconds")
	}
	if flags&log.Llongfile == log.Llongfile {
		buffer.WriteString("|Llongfile")
	}
	if flags&log.Lshortfile == log.Lshortfile {
		buffer.WriteString("|Lshortfile")
	}
	if flags&log.LUTC == log.LUTC {
		buffer.WriteString("|Ldate")
	}
	if flags&log.Lmsgprefix == log.Lmsgprefix {
		buffer.WriteString("|Lmsgprefix")
	}
	if flags&log.LstdFlags == log.LstdFlags {
		buffer.WriteString("|LstdFlags")
	}

	if buffer.Len() <= 0 {
		return ""
	}
	return buffer.String()[1:]
}
