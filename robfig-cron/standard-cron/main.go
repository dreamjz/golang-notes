package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	c.AddFunc("* * * * * ", func() {
		log.Println("tick every 1 minute")
	})

	c.AddFunc("*/2 * * * *", func() {
		log.Println("tick every 2 minutes")
	})

	c.Start()
	time.Sleep(5 * time.Second)

	time.Sleep(5 * time.Minute)
}
