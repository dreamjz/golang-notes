package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"runtime"
)

func main() {
	homedir.DisableCache = false

	dir, _ := homedir.Dir()
	fmt.Println("Home:", dir)
	fmt.Println("OS:", runtime.GOOS)
}
