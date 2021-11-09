package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name:", os.Getenv("name"))
	fmt.Println("id:", os.Getenv("id"))
}
