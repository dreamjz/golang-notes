package main

import (
	"flag"
	"fmt"
)

var (
	logLevel string
)

const (
	defaultLevel = "debug"
	usage        = "set log level value"
)

func init() {
	flag.StringVar(&logLevel, "logLevel", defaultLevel, usage)
	flag.StringVar(&logLevel, "l", defaultLevel, usage+"(shorthand)")
}

func main() {
	flag.Parse()

	fmt.Println("Log level:", logLevel)
}
