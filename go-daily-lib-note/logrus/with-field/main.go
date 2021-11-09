package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.WithFields(logrus.Fields{
		"id":   1,
		"name": "kesa",
	}).Info("Info message")

	loggerInFunction()
}

func loggerInFunction() {
	requestLogger := logrus.WithFields(logrus.Fields{
		"user_id": 11001,
		"ip":      "192.168.2.231",
	})

	requestLogger.Info("info msg")
	requestLogger.Error("error msg")
}
