package main

import (
	logredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func init() {
	hookConfig := logredis.HookConfig{
		Host:     "localhost",
		Key:      "logKey",
		Format:   "v0",
		App:      "logrus_note",
		Hostname: "localhost",
		TTL:      3600,
		Port:     6379,
	}

	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error :%q", err)
	}
}

func main() {
	logrus.Info("info message ")

	logrus.WithField("app_name", "logrus_note").Info("additional fields")

	// do not sent log to default writer
	logrus.SetOutput(ioutil.Discard)

	logrus.Info("send to redis")
}
