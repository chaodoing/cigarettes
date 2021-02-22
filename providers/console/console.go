package console

import (
	"os"

	"github.com/urfave/cli"
)

func Console() {
	app := cli.NewApp()
	app.Name = os.Getenv("APP_NAME")
	app.Usage = "排班管理系统"
	app.Version = os.Getenv("VERSION")
	app.Commands = []cli.Command{
		Servcie,
		Systemd,
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
