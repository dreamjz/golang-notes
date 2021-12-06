package main

import (
	"fmt"
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
	c := cron.New(cron.WithChain(wrapper1(), wrapper2(), wrapper3()))
	c.AddJob("@every 1s", HelloJob{"kesa"})
	c.Start()
	time.Sleep(5 * time.Second)
}

func wrapper1() cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			fmt.Println("w1 before")
			j.Run()
			fmt.Println("w1 after")
		})
	}
}

func wrapper2() cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			fmt.Println("w2 before")
			j.Run()
			fmt.Println("w2 after")
		})
	}
}

func wrapper3() cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			fmt.Println("w3 before")
			j.Run()
			fmt.Println("w3 after")
		})
	}
}
