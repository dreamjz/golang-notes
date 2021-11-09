package main

import (
	"flag"
	"fmt"
)

var (
	intFlag    int
	boolFlag   bool
	stringFlag string
)

func init() {
	flag.IntVar(&intFlag, "intFlag", 0, "int flag value")
	flag.BoolVar(&boolFlag, "boolFlag", false, "bool flag value")
	flag.StringVar(&stringFlag, "stringFlag", "default", "string flag value")
}

func main() {
	flag.Parse()

	fmt.Println("Non-flag command-line arguments:", flag.Args())
	fmt.Println("The number of non-flag command-line arguments:", flag.NArg())
	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("%d'th argument remaining after flags have been processed:%s\n", i, flag.Arg(i))
	}
	fmt.Println("The number of flags have been set:", flag.NFlag())

	fmt.Println("int flag:", intFlag)
	fmt.Println("bool flag:", boolFlag)
	fmt.Println("string flag:", stringFlag)
}
