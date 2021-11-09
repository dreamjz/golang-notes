package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"log"
)

func main() {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Home dir:", dir)

	dir = "~/dir"
	expandedDir, err := homedir.Expand(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Expand of %s is :%s\n", dir, expandedDir)
}
