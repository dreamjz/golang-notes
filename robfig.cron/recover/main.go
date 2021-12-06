package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type PanicJob struct{}

func (PanicJob) Run() {
	panic("job panic")
}
func main() {
	c := cron.New()
	c.AddJob("@every 1s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(PanicJob{}))
	c.Start()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
