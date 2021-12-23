package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
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
	file, err := createLogFile("./log/cron.log")
	if err != nil {
		log.Fatal("create file err: ", err.Error())
	}
	writer := io.MultiWriter(file, os.Stdout)
	logger := cron.VerbosePrintfLogger(log.New(writer, "[CRON]: ", log.LstdFlags))
	c := cron.New(cron.WithLogger(logger))
	c.AddJob("@every 1s", HelloJob{"kesa"})
	c.Start()
	time.Sleep(5 * time.Second)
}

func createLogFile(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	log.Println("Log dir: ", dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal("create log dir err: ", err.Error())
		}
	}
	return os.Create(path)
}
