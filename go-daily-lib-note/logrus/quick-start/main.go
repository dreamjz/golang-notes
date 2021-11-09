package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("Something very low level.")
	logrus.Debug("Useful debugging information.")
	logrus.Info("Something noteworthy happened!")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	logrus.Fatal("Bye.")
	// Calls panic() after logging
	logrus.Panic("I'm bailing.")
}
