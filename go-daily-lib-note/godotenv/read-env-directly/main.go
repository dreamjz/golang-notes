package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	envStr := `
# app name
app_name = godotenv_note
# app version
version = 0.0.1
`
	myEnv, err := godotenv.Unmarshal(envStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("App name:", myEnv["app_name"])
	fmt.Println("version:", myEnv["version"])
}
