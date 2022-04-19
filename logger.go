package muma

import (
	"os"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	ExitsFile("log")

	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.Level(zapcore.DebugLevel)),
		Development:      config.Debug,
		Encoding:         "json",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"log/debug.log"},
		ErrorOutputPaths: []string{"stderr", "log/error.log"},
	}

	if runtime.GOOS == "windows" {
		cfg.Encoding = "console"
		cfg.OutputPaths = append(cfg.OutputPaths, "stdout")
	}

	var err error
	Logger, err = cfg.Build()

	if err != nil {
		panic(err)
	}
}

func ExitsFile(path string) {
	_, err := os.Stat(path)

	if os.IsExist(err) {
		return
	}

	err = os.MkdirAll("log", os.ModeDir)

	if err != nil {
		panic(err)
	}
}
