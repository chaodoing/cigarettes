package api

import (
	"github.com/chaodoing/cigarettes/providers/container"
	"github.com/chaodoing/cigarettes/providers/utils"
	"github.com/kataras/iris/v12"
)

func Index(ctx iris.Context, container container.Container)  {
	utils.Export(ctx, 0, "Ok", iris.Map{"path": "Index"})
}