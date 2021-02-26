package api

import (
	"github.com/chaodoing/cigarettes/providers/container"
	"github.com/chaodoing/cigarettes/providers/utils"
	"github.com/kataras/iris/v12"
	"strings"
)

// Export 导出数据
//	@param iris.Context ctx
//	@param container.Containers container
func Export(ctx iris.Context, container container.Containers) {
	uid := ctx.URLParamDefault("uuid", "")
	if strings.EqualFold(uid, "") {
		utils.Export(ctx, 1, "您需要查询的uuid错误", nil)
		return
	} else {

	}
}
