package logger

import (
	"demo/pkg/tools"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// Logger  is a globle
var Logger *zap.Logger

type LogConfig struct {
	Dir string `yaml:"dir" default:"log" `
	Level  zapcore.Level `yaml:"level" default:"debug"`
	MaxFile int  `default:"7"`
	MaxAge int `default:"1"`
}


func InitLogger(cfg  LogConfig)  {
	app := tools.AppName()
	file:=  fmt.Sprintf("%s/%s.log",cfg.Dir , app )//filePath
	hook := lumberjack.Logger{
		Filename: file,
		MaxBackups: cfg.MaxFile,
		MaxAge:     cfg.MaxAge,     //days
		Compress:   true, // disabled by default
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "l",
		NameKey:        "logger",
		CallerKey:      "file",
		MessageKey:     "m",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     func (t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			type appendTimeEncoder interface {
				AppendTimeLayout(time.Time, string)
			}
			if enc, ok := enc.(appendTimeEncoder); ok {
				enc.AppendTimeLayout(t, "2006-01-02 15:04:05.000")
				return
			}
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >=  cfg.Level
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		//zapcore.AddSync(&hook), // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		level)

	Logger = zap.New(core, zap.AddCaller())

	//Logger, _ = zap.NewProduction()

	defer func() {
		_ = Logger.Sync()
	}()
	Logger.Info("logger start")

}

func Info(msg string, fields ...zap.Field)  {
	if Logger != nil {
		Logger.Info(msg,fields...)
	}
}

func Debug(msg string, fields ...zap.Field)  {
	if Logger != nil {
		Logger.Debug(msg,fields...)
	}
}

func Warn(msg string, fields ...zap.Field)  {
	if Logger != nil {
		Logger.Warn(msg,fields...)
	}
}
func Error(msg string, fields ...zap.Field)  {
	if Logger != nil {
		Logger.Error(msg,fields...)
	}
}