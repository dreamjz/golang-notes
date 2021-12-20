package main

import (
	"os"

	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
)

func main() {
	consoleWriter := zapcore.Lock(os.Stdout)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		consoleWriter, zapcore.DebugLevel)
	logger := zap.New(core)
	defer logger.Sync()
	logger.Info("advanced configuration")
}
