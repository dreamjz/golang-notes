package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type HelloJob struct {
	Name string
}

func (h HelloJob) Run() {
	log.Println("Hello", h.Name)
}

func main() {
	c := cron.New()
	c.AddJob("@every 1s", HelloJob{"kesa"})
	c.Start()
	time.Sleep(5 * time.Second)
}
