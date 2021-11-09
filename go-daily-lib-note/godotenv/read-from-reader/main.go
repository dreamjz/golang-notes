package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	file, _ := os.OpenFile(".env", os.O_RDONLY, 0666)
	myEnv, err := godotenv.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("app name:", myEnv["app_name"])
	fmt.Println("version:", myEnv["version"])

	buf := &bytes.Buffer{}
	buf.WriteString("app_name: godotenv_note@buffer")
	buf.WriteString("\n")
	buf.WriteString("version: 0.0.1")
	bufEnv, err := godotenv.Parse(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("app name:", bufEnv["app_name"])
	fmt.Println("version:", bufEnv["version"])
}
