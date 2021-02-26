package route

import (
	"github.com/chaodoing/cigarettes/modules/api"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func Initialize(app *iris.Application) {
	root := app.Party(`/`)
	{
		root.Get(`/`, hero.Handler(api.Index))
		root.Get(`/api`, hero.Handler(api.Index))
		root.Get(`/api/index`, hero.Handler(api.Index))
	}
	API := app.Party(`/api`)
	{
		// 输入数据
		API.Post(`/input`, hero.Handler(api.Input))
		// 输出数据
		API.Get(`/export`, hero.Handler(api.Export))
	}
}
