package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

var FlagHasBeenSet error = errors.New("flag has been set")

type interval []time.Duration

func (i interval) String() string {
	return fmt.Sprint(([]time.Duration)(i))
}

func (i *interval) Set(value string) error {
	if len(*i) > 0 {
		return FlagHasBeenSet
	}
	for _, v := range strings.Split(value, ",") {
		d, err := time.ParseDuration(v)
		if err != nil {
			return err
		}
		*i = append(*i, d)
	}
	return nil
}

var intervalFlag interval

func init() {
	flag.Var(&intervalFlag, "deltaT", "comma-seperated list of intervals to use between events ")
}

func main() {
	flag.Parse()

	fmt.Println(intervalFlag)
}
