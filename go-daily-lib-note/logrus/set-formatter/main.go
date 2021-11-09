package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Info message")
}
