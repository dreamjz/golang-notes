package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := godotenv.Load("common", "dev.env", "production.env")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("App name:", os.Getenv("app_name"))
	fmt.Println("Version::", os.Getenv("version"))
	fmt.Println("Database:", os.Getenv("database"))

}
