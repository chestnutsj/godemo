package main

import (
	"demo/pkg/config"
	"demo/pkg/logger"
	"demo/pkg/tools"
	"flag"
	"fmt"
	"os"
	"runtime"
)
var (
	// 初始化为 unknown，如果编译时没有传入这些值，则为 unknown
	GitCommitLog   = "unknown"
	GitStatus      = "unknown"
	BuildTime      = "unknown"
	BuildGoVersion = "unknown"
)


func main()  {

	v := flag.Bool("v", false, "show bin info")
	conf := flag.String("c","conf/"+tools.AppName()+".yaml", "show bin info")
	flag.Parse()
	if *v {
		fmt.Printf("GitCommitLog=%s\nGitStatus=%s\nBuildTime=%s\nGoVersion=%s\nruntime=%s/%s\n",
			GitCommitLog, GitStatus, BuildTime, BuildGoVersion, runtime.GOOS, runtime.GOARCH)
		return
	}

	cfg,err := config.InitConfig(*conf)
	if err != nil {
		os.Exit(1)
	}
 	logger.InitLogger(cfg.Log)
	logger.Info("start app "+ tools.AppName())

}
