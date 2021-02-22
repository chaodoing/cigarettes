package utils

import (
	"github.com/kataras/iris/v12"

	"github.com/chaodoing/cigarettes/providers/container"
)

// Export 输出数据
//	@param iris.Context ctx
//	@param int Code 错误代码
//	@param string Message 错误消息
//	@param interface data 数据
//	@return (int, err)
func Export(ctx iris.Context, Code int, Message string, Data interface{}) (int, error) {
	var response = &container.Response{}
	return response.Context(ctx).Data(Code, Message, Data).Send()
}

// Page 输出分页数据
//	@param iris.Context ctx
//	@param int Code 错误代码
//	@param string Message 错误消息
//	@param interface data 数据
//	@return (int, err)
func Page(ctx iris.Context, page container.Page) (int, error) {
	var response = &container.Response{}
	return response.Context(ctx).SetPage(page).Send()
}
