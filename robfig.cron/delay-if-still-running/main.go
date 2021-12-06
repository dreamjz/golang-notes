package main

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type DelayJob struct {
	count int
}

func (d *DelayJob) Run() {
	time.Sleep(2 * time.Second)
	d.count++
	log.Printf("%d delay job", d.count)
}

func main() {
	c := cron.New()
	c.AddJob("@every 1s", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&DelayJob{}))
	c.Start()
	time.Sleep(5 * time.Second)
	fmt.Println("main end")
}
