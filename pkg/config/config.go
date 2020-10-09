package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chestnutsj/godemo/pkg/logger"
	"github.com/jinzhu/configor"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Port  int
	Metric int `default:"9100"`
 	Log logger.LogConfig `yaml:"log"`
	DbUrl string
	Remote string
	cfgFile string `yaml:"-" json:"-"`
}

func InitConfig(file string,flush bool ) (*Config, error)  {
	cfg:=Config{}
	err:=  configor.New(&configor.Config{AutoReload: flush, AutoReloadCallback: func(config interface{}) {
		fmt.Printf("%v changed", config)
	}}).Load(&cfg, file)
	if err != nil {
		logger.Error("read config failed"+ err.Error())
		return nil,err
	}
	cfg.cfgFile = file
	return  &cfg,nil
}


func (c *Config) Save() {
	data,err := yaml.Marshal(c)
	if err != nil {
	 	logger.Error(err.Error())
		panic(err)
	}
	err = ioutil.WriteFile(c.cfgFile,data,0777)
	if err != nil {
		logger.Error(err.Error())
	}
}

func (conf *Config) String() string {
	b, err := json.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	return out.String()
}