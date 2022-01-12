package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

func main() {
	// create trace file
	f, err := os.OpenFile("trace.out", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("failed to create trace file: %v", err)
	}
	// remember to close
	defer f.Close()

	// start trace goroutine
	err = trace.Start(f)
	if err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	// remember to stop
	defer trace.Stop()

	// main
	fmt.Println("Hello World")
}
