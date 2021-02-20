package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)
var (
	stdCancel  func()
	replaceCancel  func()
	once sync.Once
	once2 sync.Once
	Logger *zap.Logger
	level zap.AtomicLevel
	stdLevel zap.AtomicLevel
)


type LogConfig struct {

	Dir     string `yaml:"dir" default:"log" `
	Level   string `yaml:"level" default:"info"`
	MaxFile int    `default:"7"`
	MaxAge  int    `default:"9"`
	Compress bool  `default:"true"`
	Std     bool `default:"true"`
}

func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}

func SetStd(b bool)  {
	if b {
		stdLevel.SetLevel(zap.DebugLevel)
	} else {
		stdLevel.SetLevel(zap.PanicLevel)
	}
}

func SetLevel(l string)  {
	_ = level.UnmarshalText([]byte(l))
}

func init()  {
	level = zap.NewAtomicLevel()
	stdLevel = zap.NewAtomicLevel()
	once.Do(func() {
		Logger, _ = zap.NewDevelopment()
		stdCancel = zap.RedirectStdLog(Logger)
		replaceCancel = zap.ReplaceGlobals(Logger)
		Sync()
	})
}

func GetExeName() string  {
	exe,err:=  os.Executable()
	if err != nil {
		log.Print("get Executable failed")
		exe=os.Args[0]
	}
	_,f:= filepath.Split(exe)
	return  f
}

func InitConfig(cfg LogConfig)  {
	once2.Do(func() {


	file := fmt.Sprintf("%s/%s.log",cfg.Dir,GetExeName()) //filePath
	hook := lumberjack.Logger{
			Filename:   file,
			MaxBackups: cfg.MaxFile,
			MaxAge:     cfg.MaxAge, //days
			Compress:   cfg.Compress,       // disabled by default
			}

	encoderConfig := zapcore.EncoderConfig{
		FunctionKey:    zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LevelKey:      "lv",
		TimeKey:       "ts",
		CallerKey:      "line",
		EncodeLevel:   zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
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
		LineEnding:     zapcore.DefaultLineEnding,

		EncodeCaller:   zapcore.FullCallerEncoder,

	}
	SetLevel(cfg.Level)

	SetStd(cfg.Std)

	core:= zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(&hook),
			level),
		zapcore.NewCore( zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()) ,
			zapcore.AddSync(os.Stdout),
			stdLevel),
		)

	Logger = zap.New(core,zap.AddCaller())
	if stdCancel != nil {
		stdCancel()
	}
	stdCancel = zap.RedirectStdLog(Logger)

	if replaceCancel != nil {
		replaceCancel()
	}
	replaceCancel = zap.ReplaceGlobals(Logger)
	defer func() {
		_ = Logger.Sync()
	}()
	Logger.Info("logger start")
	})
}
