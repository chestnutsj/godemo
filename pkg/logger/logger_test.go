package logger

import (
	"go.uber.org/zap/zapcore"
	"testing"
)


func Test_initLogger(t *testing.T) {

	cfg:=LogConfig{
		Dir:"log",
		Level: zapcore.DebugLevel,
		MaxFile: 7,
		MaxAge: 1,
	}

	InitLogger(cfg)
	Debug("DDD")
	Info("aaa")
	Warn("www")
	Error("eeee")
}
