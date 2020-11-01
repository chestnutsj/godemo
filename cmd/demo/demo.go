package main

import (
	"flag"
	"fmt"
	"github.com/chestnutsj/godemo/pkg/config"
	"github.com/chestnutsj/godemo/pkg/logger"
	"github.com/chestnutsj/godemo/pkg/tools"
	"path/filepath"
	"runtime"
)

var (
	// 初始化为 unknown，如果编译时没有传入这些值，则为 unknown
	GitCommitLog   = "unknown"
	GitStatus      = "unknown"
	BuildTime      = "unknown"
	BuildGoVersion = "unknown"
)

func main() {

	v := flag.Bool("v", false, "show bin info")
	conf := flag.String("c", filepath.Join("conf", tools.AppName()+".yaml"), "show bin info")
	flush := flag.Bool("f", false, " auto flush config")

	flag.Parse()
	if *v {
		fmt.Printf("GitCommitLog=%s\nGitStatus=%s\nBuildTime=%s\nGoVersion=%s\nruntime=%s/%s\n",
			GitCommitLog, GitStatus, BuildTime, BuildGoVersion, runtime.GOOS, runtime.GOARCH)
		return
	}

	absFile, err := filepath.Abs(*conf)
	if err != nil {
		panic(err)
	}

	cfg, err := config.InitConfig(absFile, *flush)
	if err != nil {
		panic(err)
	}
	logger.InitLogger(cfg.Log)
	logger.Info("start app " + tools.AppName())
	defer logger.Sync()
	fmt.Println(cfg.String())
	cfg.Save()

}
