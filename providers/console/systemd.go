package console

import (
	"html/template"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var Systemd = cli.Command{
	Name:        "systemd",
	ShortName:   "system",
	Usage:       "创建服务脚本",
	Description: "使用 COMMAND_NAME systemd 创建服务脚本",
	Action:      runSystemd,
	Flags: []cli.Flag{
		StringFlag("appname, a", "", "脚本输出位置"),
		StringFlag("directory, d", "", "脚本输出位置"),
		StringFlag("exec, e", "", "脚本输出位置"),
	},
}

func runSystemd(c *cli.Context) {
	var (
		appName   = os.Getenv("APP_NAME")
		directory = "/opt/" + os.Getenv("APP_NAME")
		exec      = "/opt/" + os.Getenv("APP_NAME") + "/bin/" + os.Getenv("APP_NAME") + " web"
	)
	if !strings.EqualFold(c.String("app_name"), "") {
		appName = c.String("app_name")
	}
	if !strings.EqualFold(c.String("directory"), "") {
		directory = c.String("directory")
	}
	if !strings.EqualFold(c.String("exec"), "") {
		exec = c.String("exec")
	}
	dat := data{
		Name: appName,
		Dir:  directory,
		Exec: exec,
	}
	tpl, err := template.New("systemd").Parse(systemd)
	if err != nil {
		panic(err)
		return
	}
	err = tpl.Execute(os.Stdout, dat)
	if err != nil {
		panic(err)
		return
	}
}
