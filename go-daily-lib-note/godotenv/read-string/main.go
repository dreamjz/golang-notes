package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	content := `
app_name: godotenv_note@str
version: 0.0.1
`
	strEnv, err := godotenv.Unmarshal(content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("App name:", strEnv["app_name"])
	fmt.Println("version:", strEnv["version"])
}
