package main

import (
	"fmt"
	"log"
	"time"

	"github.com/uniplaces/carbon"
)

func main() {
	outputFormat := "%-30s:%s\n"
	// calculate date with 'time'
	fmt.Println("Calculate date with 'time'")
	now := time.Now()

	fmt.Printf(outputFormat, "now", now)
	fmt.Printf(outputFormat, "one second later", now.Add(time.Second))
	fmt.Printf(outputFormat, "one minute later", now.Add(time.Minute))
	fmt.Printf(outputFormat, "one hour later", now.Add(time.Hour))

	dur, err := time.ParseDuration("3m20s")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(outputFormat, "3 minutes and 20 seconds later", now.Add(dur))

	dur, err = time.ParseDuration("2h30m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(outputFormat, "2 hours and 30 minutes later", now.Add(dur))
	// call AddDate instead of Add
	fmt.Printf(outputFormat, "3 days and 2 hours later", now.AddDate(0, 0, 3).Add(time.Hour*2))

	// calculate date with 'uniplace-carbon'
	fmt.Println("Calculate date with 'uniplace-carbon'")
	cNow := carbon.Now()

	fmt.Printf(outputFormat, "now", cNow)
	fmt.Printf(outputFormat, "one second later", cNow.AddSecond())
	fmt.Printf(outputFormat, "one minute later", cNow.AddMinute())
	fmt.Printf(outputFormat, "one hour later", cNow.AddHour())
	fmt.Printf(outputFormat, "3 minutes and 20 seconds later", cNow.AddMinutes(3).AddSeconds(20))
	fmt.Printf(outputFormat, "2 hours and 30 minutes later", cNow.AddHours(2).AddMinutes(30))
	fmt.Printf(outputFormat, "3 days and 2 hours later", cNow.AddDays(3).AddHours(2))
}
