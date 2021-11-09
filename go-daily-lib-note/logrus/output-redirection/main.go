package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

func main() {
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("./output_redirection.log", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("create file failed :", err)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	logrus.SetOutput(io.MultiWriter(writer3, writer2, writer1))
	logrus.Info("Info message")
	fmt.Println("BUF:", writer1.String())
}
