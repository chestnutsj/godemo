package config

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"log"

	"github.com/chestnutsj/godemo/pkg/tools"
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"syscall"
)

var Cfg *Config

type Config struct {
	Metric string    `default:"0"`
	Log    LogConfig `yaml:"log"`
}

type LogConfig struct {
	Std     bool          `yaml:"std" default:"true"`
	Dir     string        `yaml:"dir" default:"log" `
	Level   zapcore.Level `yaml:"level" default:"info"`
	MaxFile int           `default:"7"`
	MaxAge  int           `default:"1"`
}

func ConfigInit() {
	configFile := getConfigFile()

	_, flush := syscall.Getenv("CFG_FILE_AUTO_RELOAD")
	cfg := Config{}
	err := configor.New(&configor.Config{AutoReload: flush, AutoReloadCallback: func(config interface{}) {
		fmt.Printf("%v changed", config)
	}}).Load(&cfg, configFile)
	if err != nil {
		log.Panic("can't read config file", zap.String("fileName", configFile), zap.Error(err))
	}
	Cfg = &cfg
	log.Println(cfg.String())
}

func (c *Config) Save() {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(getConfigFile(), data, 0777)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (conf *Config) String() string {
	b, err := yaml.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}

	return string(b)
}

func getConfigFile() string {
	prefix := os.Getenv("CONFIGOR_ENV_PREFIX")
	if len(prefix) > 0 {
		prefix += "_"
	}
	env_cfg := prefix + "CFG_FILE"
	configFile, exists := syscall.Getenv(env_cfg)
	if !exists {
		configFile = "conf/" + tools.AppName() + ".yml"
	}
	return configFile
}
