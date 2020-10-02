package main

import (
	"flag"
	"mwxy_create_articles/bootstrap"
	"mwxy_create_articles/pkg/conf"
	"mwxy_create_articles/pkg/util"
	"mwxy_create_articles/routers"
	"os/exec"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "c", util.RelativePath("conf.ini"), "配置文件")
	flag.Parse()
	// bootstrap
	bootstrap.Init(confPath)
}

func main() {
	api := routers.InitRouter()
	util.Log().Info("开始监听 %s", conf.SystemConfig.Listen)
	//打开浏览器
	exec.Command(`cmd`, `/c`, `start`, `http://127.0.0.1`+conf.SystemConfig.Listen).Start()

	if err := api.Run(conf.SystemConfig.Listen); err != nil {
		util.Log().Error("无法监听[%s]，%s", conf.SystemConfig.Listen, err)
	}
}
