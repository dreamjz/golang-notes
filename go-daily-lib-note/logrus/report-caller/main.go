package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetReportCaller(true)

	logrus.Info("Something noteworthy happened!")
}
