package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	buf := &bytes.Buffer{}
	buf.WriteString("app: @buf\n")
	buf.WriteString("ver: 0.1")
	bufEnv, _ := godotenv.Parse(buf)

	content, err := godotenv.Marshal(bufEnv)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("String env: ", content)
}
