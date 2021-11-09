package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("app", os.Getenv("app"))
	fmt.Println("ver", os.Getenv("ver"))
}
