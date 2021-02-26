package main

import (
	"github.com/chaodoing/cigarettes/providers/console"
	"github.com/chaodoing/cigarettes/providers/container"
	"github.com/chaodoing/cigarettes/providers/service"
	"os"
	"strings"
)

var (
	ENV     string // 环境
	VERSION string // 版本
	NAME    string // App名称
)

// initialize 初始化参数
func initialize() {
	os.Setenv("APP_ENV", ENV)
	os.Setenv("APP_VERSION", VERSION)
	os.Setenv("APP_NAME", NAME)
}

// main 主函数
func main() {
	initialize()

	if strings.EqualFold(ENV, "release") {
		console.Console()
	} else {
		var (
			env        = container.Env(os.Getenv("PWD") + "/conf/app.ini")
			containers = container.XContainer(env)
		)
		service.AppService(containers)
	}
}
