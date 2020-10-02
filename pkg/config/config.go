package config

import (
	"demo/pkg/logger"
	"fmt"
	"github.com/jinzhu/configor"
)

type Config struct {
	Port  int
 	Log logger.LogConfig `yaml:"log"`
	DbUrl string
	Remote string
}

var cfg Config

func InitConfig(file string ) (*Config, error)  {
	err:=  configor.New(&configor.Config{AutoReload: true, AutoReloadCallback: func(config interface{}) {
		fmt.Printf("%v changed", config)
	}}).Load(&cfg, file)
	if err != nil {
		fmt.Println("read config failed"+ err.Error())
		return nil,err
	}
	return  &cfg,nil
}
