package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

type Option struct {
	IntFlag        int            `short:"i" long:"int" description:"int flag value"`
	IntSlice       []int          `long:"intSlice" description:"int slice flag value"`
	BoolFlag       bool           `long:"bool" description:"bool flag value"`
	BoolSlice      []bool         `long:"boolSlice" description:"bool slice flag value"`
	FloatFlag      float64        `long:"float" description:"float flag value"`
	FloatSlice     []float64      `long:"floatSlice" description:"float slice flag value"`
	StringFlag     string         `short:"s" long:"string" description:"string flag value"`
	StringSlice    []string       `long:"stringSlice" description:"string slice flag value"`
	PtrStringSlice []*string      `short:"p" long:"ptrStrSlice" description:"pointer of string slice flag value"`
	Call           func(string)   `long:"call" description:"callback"`
	IntMap         map[string]int `long:"intMap" description:"a map from string to int"`
}

func main() {
	var opt Option
	opt.Call = func(value string) {
		fmt.Println("in callback:", value)
	}
	_, err := flags.Parse(&opt)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	fmt.Println("int flag:", opt.IntFlag)
	fmt.Println("int slice:", opt.IntSlice)
	fmt.Println("bool flag:", opt.BoolFlag)
	fmt.Println("bool slice:", opt.BoolSlice)
	fmt.Println("float flag:", opt.FloatFlag)
	fmt.Println("float slice:", opt.FloatSlice)
	fmt.Println("string flag:", opt.StringFlag)
	fmt.Println("string slice:", opt.StringSlice)
	fmt.Println("pointer of string slice:")
	for i, v := range opt.PtrStringSlice {
		fmt.Printf("\t%d: %v\n", i, *v)
	}
	fmt.Println("int map:", opt.IntMap)
}
