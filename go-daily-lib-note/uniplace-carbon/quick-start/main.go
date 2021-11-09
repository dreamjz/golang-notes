package main

import (
	"fmt"
	"github.com/uniplaces/carbon"
	"time"
)

func main() {
	fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())

	today, _ := carbon.NowInLocation("Japan")
	fmt.Printf("Right now in Japan is %s\n", today)

	fmt.Printf("Tomorrow is %s\n", carbon.Now().AddDay())
	fmt.Printf("Last week is %s\n", carbon.Now().SubWeek())

	nextOlympics, _ := carbon.CreateFromDate(2016, time.August, 5, "Europe/London")
	nextOlympics = nextOlympics.AddYears(4)
	fmt.Printf("Next Olympics are in %d\n", nextOlympics.Year())

	if carbon.Now().IsWeekend() {
		fmt.Printf("Happy time")
	}
}
