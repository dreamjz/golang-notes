package main

import (
	"fmt"
	"log"
	"time"

	"github.com/uniplaces/carbon"
)

func main() {
	date, err := carbon.Create(2021, 11, 6, 0, 0, 0, 0, "Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}
	date.SetWeekStartsAt(time.Sunday)
	date.SetWeekEndsAt(time.Saturday)
	date.SetWeekendDays([]time.Weekday{time.Monday, time.Tuesday, time.Thursday, time.Friday})

	fmt.Printf("Today is %s,weekend? %t\n", date.Weekday(), date.IsWeekend())
}
