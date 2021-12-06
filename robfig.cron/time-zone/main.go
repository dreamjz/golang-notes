package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// set loc
	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal(err.Error())
	}
	c := cron.New(cron.WithLocation(tz))
	c.AddFunc("0 6 * * *", func() {
		log.Println("Every 6 o'clock in Shanghai")
	})
	c.AddFunc("CRON_TZ=Asia/Tokyo 0 6 * * *", func() {
		log.Println("Every 6 o'clock in Tokyo")
	})
	c.Start()
}
