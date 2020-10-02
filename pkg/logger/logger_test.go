package logger

import (
	"github.com/chestnutsj/godemo/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func Test_initLogger(t *testing.T) {

	cfg := config.LogConfig{
		Dir:     "log",
		Level:   zapcore.DebugLevel,
		MaxFile: 7,
		MaxAge:  1,
	}

	InitLogger(cfg)
	zap.L().Info(("Info"))
	zap.L().Debug(("Debug"))
	zap.L().Warn(("Warn"))
	zap.L().Error(("Error"))
}
