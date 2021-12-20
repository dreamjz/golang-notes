package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const num = 10000

func main() {
	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/foo.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	consoleSyncer := zapcore.Lock(os.Stdout)

	fileEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	//syncer := zapcore.NewMultiWriteSyncer(w)
	//core := zapcore.NewCore(
	//	zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
	//	syncer,
	//	zap.InfoLevel,
	//)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, w, zap.InfoLevel),
		zapcore.NewCore(consoleEncoder, consoleSyncer, zap.InfoLevel),
	)
	logger := zap.New(core)
	for i := 0; i < num; i++ {
		logger.Info("log rotation: ", zap.Int("No", i))
	}
}
