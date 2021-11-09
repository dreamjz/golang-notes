package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	period time.Duration
)

func init() {
	flag.DurationVar(&period, "period", 1*time.Second, "set sleep period")
}

func main() {
	flag.Parse()

	fmt.Printf("Sleeping for %v\n", period)
	time.Sleep(period)
	fmt.Println("Wake up ...")
}
