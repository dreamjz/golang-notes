package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

type SkipJob struct {
	count int64
}

func (s *SkipJob) Run() {
	s.count++
	log.Printf("%d job ", s.count)
	if s.count == 1 {
		time.Sleep(2 * time.Second)
	}
}

func main() {
	c := cron.New()
	logger := cron.VerbosePrintfLogger(log.New(os.Stdout, "[CRON]: ", log.LstdFlags))
	c.AddJob("@every 1s", cron.NewChain(cron.SkipIfStillRunning(logger)).Then(&SkipJob{}))
	c.Start()
	time.Sleep(10 * time.Second)
	fmt.Println("main end")
}
