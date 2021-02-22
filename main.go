package main

import (
	"github.com/chaodoing/cigarettes/providers/container"
	"github.com/chaodoing/cigarettes/providers/console"
	"github.com/chaodoing/cigarettes/providers/service"
	"os"
	"strings"
)

var (
	ENV     string	// 环境
	VERSION string	// 版本
	NAME    string  // App名称
)

func initialize() {
	os.Setenv("APP_ENV", ENV)
	os.Setenv("APP_VERSION", VERSION)
	os.Setenv("APP_NAME", NAME)
}
func main() {
	initialize()
	if strings.EqualFold(ENV, "release") {
		console.Console()
	} else {
		var (
			config = os.Getenv("PWD") + "/conf/app.ini"
			env    = container.Env(config)
			containers := container.XContainer(env)
			service.AppService(containers)
		)
	}
}