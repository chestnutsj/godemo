package main

import (
	"flag"
	"fmt"
	"github.com/chestnutsj/godemo/pkg/config"
	"github.com/chestnutsj/godemo/pkg/logger"
	"github.com/chestnutsj/godemo/pkg/metrics"
	"log"
	"os"

	"runtime"
)

var (
	// 初始化为 unknown，如果编译时没有传入这些值，则为 unknown
	GitCommitLog   = "unknown"
	BuildTime      = "unknown"
	BuildGoVersion = "unknown"
)

func main() {
	v := flag.Bool("version", false, "show bin info")
	genCfg := flag.Bool("genCfg", false, "show config file")
	flag.Parse()
	if *v {
		fmt.Printf("GitCommitLog=%s\nBuildTime=%s\nGoVersion=%s\nruntime=%s/%s\n",
			GitCommitLog, BuildTime, BuildGoVersion, runtime.GOOS, runtime.GOARCH)
		return
	}

	config.ConfigInit()
	if config.Cfg == nil {
		log.Panic("config is not set")
	}
	if *genCfg {
		println(config.Cfg.String())
		config.Cfg.Save()
		os.Exit(0)
	}

	logger.InitLogger(config.Cfg.Log)

	config.Cfg.Save()
	metrics.StartMetrics(config.Cfg.Metric)
}
