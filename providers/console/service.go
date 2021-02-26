package console

import (
	"os"
	"strings"

	"github.com/urfave/cli"

	"github.com/chaodoing/cigarettes/providers/container"
	"github.com/chaodoing/cigarettes/providers/service"
)

var Servcie = cli.Command{
	Name:        "service",
	ShortName:   "web",
	Usage:       "启动网站服务",
	Description: "使用 COMMAND_NAME service 启动网站服务",
	Action:      runService,
	Flags: []cli.Flag{
		StringFlag("port, p", "", "使用监听端口"),
		StringFlag("config, c", os.Getenv("PWD")+"/conf/app.ini", "自定义配置文件路径"),
	},
}

// runService
func runService(c *cli.Context) {
	config := c.String("config")
	environment := container.Env(config)
	port := c.String("port")
	if !strings.EqualFold(port, "") {
		environment.App.Port = port
	}
	var containers = container.XContainer(environment)
	service.AppService(containers)
}
