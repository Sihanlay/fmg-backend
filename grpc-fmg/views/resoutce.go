package views

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"grpc-demo/views/resource"
)

func RegisterResourceRouters(app *iris.Application) {

	// 资源路由
	resourceRouter := app.Party("/resources")

	resourceRouter.Get("/qiniu/upload_token", hero.Handler(resource.GetQiNiuUploadToken))
}
