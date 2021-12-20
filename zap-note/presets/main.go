package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	url := "example"
	devLogger, _ := zap.NewDevelopment()
	defer devLogger.Sync()
	devLogger.Info("failed to fetch url",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))

	exampleLogger := zap.NewExample()
	defer exampleLogger.Sync()
	exampleLogger.Info("failed to fetch url",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))

	prodLogger, _ := zap.NewProduction()
	defer prodLogger.Sync()
	prodLogger.Info("failed to fetch url",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))

}
