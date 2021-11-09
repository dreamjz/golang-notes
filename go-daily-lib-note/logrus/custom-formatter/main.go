package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type AppHook struct {
	AppName string
}

func (ah *AppHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.InfoLevel}
}

func (ah *AppHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Fire !")
	entry.Data["app"] = ah.AppName
	return nil
}

func main() {
	myHook := &AppHook{AppName: "logrus_note"}
	logrus.AddHook(myHook)

	logrus.Info("info message")
	logrus.Warn("waring message")
}
