package main

import (
	"flag"
	"fmt"
)

var (
	intFlag    *int
	boolFlag   *bool
	stringFlag *string
)

func init() {
	intFlag = flag.Int("intFlag", 0, "int flag value")
	boolFlag = flag.Bool("boolFlag", false, "boolean flag value")
	stringFlag = flag.String("stringFlag", "default", "string flag value")
}

func main() {
	flag.Parse()

	fmt.Println("Int flag :", *intFlag)
	fmt.Println("Bool flag", *boolFlag)
	fmt.Println("String flag", *stringFlag)
}
