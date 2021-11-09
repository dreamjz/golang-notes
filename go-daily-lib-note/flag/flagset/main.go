package main

import (
	"flag"
	"fmt"
)

func main() {
	args := []string{"-intFlag", "12", "-stringFlag", "test"}

	var intFlag int
	var boolFlag bool
	var stringFlag string

	fs := flag.NewFlagSet("MyFlagSet", flag.ContinueOnError)
	fs.IntVar(&intFlag, "intFlag", 0, "set int flag value")
	fs.BoolVar(&boolFlag, "boolFlag", false, "set bool flag value")
	fs.StringVar(&stringFlag, "stringFlag", "default", "set string flag value")

	fs.Parse(args)

	fmt.Println("int flag:", intFlag)
	fmt.Println("bool flag:", boolFlag)
	fmt.Println("string flag:", stringFlag)
}
