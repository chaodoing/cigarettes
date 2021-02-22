package service

import (
	"net/http"

	"github.com/kataras/iris/v12"

	"github.com/chaodoing/cigarettes/providers/container"
)

// Crossdomain 允许跨域
func Crossdomain(ctx iris.Context, container container.Container) {
	ctx.Header("Access-Control-Allow-Origin", ctx.Request().Header.Get("Origin"))
	ctx.Header("Access-Control-Allow-Headers", "authorization, language, cache-control, content-type, if-match, if-modified-since, if-none-match, if-unmodified-since, x-requested-with")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Max-Age", "1728000")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Expose-Headers", "authorization")
	ctx.Header("Version", "1.0.0")
	ctx.Header("Author", "Neil")
	ctx.Header("Email", "chaodoing@live.com")
	ctx.Header("Site-Name", container.Env().App.Name)
	if ctx.Method() == http.MethodOptions {
		ctx.StatusCode(204)
		return
	} else {
		ctx.Next()
	}
}
