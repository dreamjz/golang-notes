package main

import (
	"bytes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// parse env
	buf := &bytes.Buffer{}
	buf.WriteString("app_name: app@env_generate\n")
	buf.WriteString("version: 0.0.1")
	bufEnv, _ := godotenv.Parse(buf)
	// save env
	err := godotenv.Write(bufEnv, "./.env")
	if err != nil {
		log.Fatal(err)
	}
}
