package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("App name:", os.Getenv("app_name"))
	fmt.Println("Version::", os.Getenv("version"))
}
