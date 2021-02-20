package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestInitConfig(t *testing.T) {
	 InitConfig(LogConfig{
		Dir:     "logs",
		MaxFile: 7,
		MaxAge:  7,
		Level: "info",
	})
	zap.L().Info("start aaaa")
}

func TestGetExeName(t *testing.T) {
	 got := GetExeName()
	 t.Log(got)
}
