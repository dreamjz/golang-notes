package main

import (
	"fmt"
	"time"

	"github.com/uniplaces/carbon"
)

func main() {
	outputFormat := "%-20s: %s\n"
	now := carbon.Now()

	fmt.Printf(outputFormat, "Start of day", now.StartOfDay())
	fmt.Printf(outputFormat, "End of day", now.EndOfDay())
	fmt.Printf(outputFormat, "Start of month", now.StartOfMonth())
	fmt.Printf(outputFormat, "End of month", now.EndOfMonth())
	fmt.Printf(outputFormat, "Start of year", now.StartOfYear())
	fmt.Printf(outputFormat, "Start of decade", now.StartOfDecade())
	fmt.Printf(outputFormat, "End of decade", now.EndOfDecade())
	fmt.Printf(outputFormat, "Start of century", now.StartOfCentury())
	fmt.Printf(outputFormat, "End of century", now.EndOfCentury())
	fmt.Printf(outputFormat, "Start of week", now.StartOfWeek())
	fmt.Printf(outputFormat, "End of week", now.EndOfWeek())
	fmt.Printf(outputFormat, "Next Wednesday", now.Next(time.Wednesday))
	fmt.Printf(outputFormat, "Previous Wednesday", now.Previous(time.Wednesday))

}
