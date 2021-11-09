package main

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

func main() {
	now := time.Now()
	timestamp := 1579615973
	timeStr := "2021-10-26 17:29:00"

	fmt.Println(cast.ToTime(now))       // 2021-10-26 17:35:55.35905014 +0800 CST m=+0.000115363
	fmt.Println(cast.ToTime(timestamp)) // 2020-01-21 22:12:53 +0800 CST
	fmt.Println(cast.ToTime(timeStr))   // 2021-10-26 17:29:00 +0000 UTC

	d, _ := time.ParseDuration("1m30s")
	ns := 30000
	strWithUnit := "130s"
	strWithoutUnit := "130"

	fmt.Println(cast.ToDuration(d))              // 1m30s
	fmt.Println(cast.ToDuration(ns))             // 30µs
	fmt.Println(cast.ToDuration(strWithUnit))    // 2m10s
	fmt.Println(cast.ToDuration(strWithoutUnit)) // 130ns
}
