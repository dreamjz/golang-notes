package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

type Option struct {
	Basic GroupBasicOption `description:"basic group" group:"basic"`
	Slice GroupSliceOption `description:"slice group" group:"slice"`
}

type GroupBasicOption struct {
	IntFlag    int     `short:"i" long:"intFlag" description:"int flag"`
	BoolFlag   bool    `short:"b" long:"boolFlag" description:"bool flag"`
	FloatFlag  float64 `short:"f" long:"floatFlag" description:"float flag"`
	StringFlag string  `short:"s" long:"stringFlag" description:"string flag"`
}

type GroupSliceOption struct {
	IntSliceFlag    []int     `long:"intSlice" description:"int slice flag"`
	BoolSliceFlag   []bool    `long:"boolSlice" description:"bool slice flag"`
	FloatSliceFlag  []float64 `long:"floatSlice" description:"float slice flag"`
	StringSliceFlag []string  `long:"stringSlice" description:"string slice flag"`
}

func main() {
	var opt Option
	p := flags.NewParser(&opt, flags.Default)
	_, err := p.ParseArgs(os.Args[1:])
	if err != nil {
		return
	}
	basicGroup := p.Command.Group.Find("basic")
	for _, option := range basicGroup.Options() {
		fmt.Printf("name:%s,value:%v\n", option.LongNameWithNamespace(), option.Value())
	}
	sliceGroup := p.Command.Group.Find("slice")
	for _, option := range sliceGroup.Options() {
		fmt.Printf("name:%s,value:%v\n", option.LongNameWithNamespace(), option.Value())
	}
}
