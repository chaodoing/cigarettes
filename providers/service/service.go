package service

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"

	"github.com/chaodoing/cigarettes/providers/container"
	"github.com/chaodoing/cigarettes/route"
)

func AppService(container *container.Containers) {
	hero.Register(container)
	var (
		env = container.WatchEnv().Env()
		app = iris.Default()
	)
	app.Logger().SetLevel(env.App.LogLevel)
	app.Logger().SetOutput(container.Logger().Writer())

	// 配置文件中开启跨域
	if env.App.Crossdomain {
		app.UseGlobal(hero.Handler(Crossdomain))
	}
	app.Favicon(env.Template.Favicon)
	// 配置模板加载目录
	app.RegisterView(iris.HTML(env.Template.Directory, env.Template.Extension))
	// 静态目录
	app.HandleDir(env.WebStatic.AccessPath, os.Getenv("PWD")+"/"+env.WebStatic.LocalPath)
	// 上传目录
	app.HandleDir(env.Upload.AccessPath, os.Getenv("PWD")+"/"+env.Upload.LocalPath)

	// 允许Options请求
	app.AllowMethods(iris.MethodOptions)
	// 调用路由
	route.Initialize(app)

	// 拼接监听地址和端口
	addr := fmt.Sprintf("%v:%v", env.App.Host, env.App.Port)
	server := &http.Server{
		Addr:              addr,
		ReadTimeout:       time.Second * 10,
		WriteTimeout:      time.Second * 30,
		ReadHeaderTimeout: time.Second * 6,
	}
	// app 配置
	app.Configure(iris.WithConfiguration(iris.Configuration{
		PostMaxMemory:       int64(env.Upload.Maximum) << 20, // 最大上传文件
		TimeFormat:          "2006-01-02 15:04:05",
		EnableOptimizations: true,
		Charset:             "UTF-8",
	}))
	// 启用监听
	if err := app.Run(iris.Server(server), iris.WithSitemap(env.App.Name)); err != nil {
		fmt.Println(err)
	}
}
