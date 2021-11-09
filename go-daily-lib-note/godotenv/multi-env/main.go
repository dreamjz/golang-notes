package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "dev"
	}
	err := godotenv.Load(".env." + appEnv)
	if err != nil {
		log.Fatal(err)
	}
	// read basic env
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("app:", os.Getenv("app"))
	fmt.Println("version:", os.Getenv("version"))
	fmt.Println("database:", os.Getenv("database"))
}
