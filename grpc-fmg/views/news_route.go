package views

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"grpc-demo/views/deliever"
)

func RegisterNewsRouters(app *iris.Application) {
	NewsRouter := app.Party("delivery/info")
	NewsRouter.Post("", hero.Handler(deliever.CreatDelivery))
}
