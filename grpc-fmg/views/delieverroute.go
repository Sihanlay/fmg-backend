package views

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"grpc-demo/views/deliever"
)

func DeliveryRouters(app *iris.Application) {

	deliveyRouter := app.Party("delivery/info")
	deliveyRouter.Post("",hero.Handler(deliever.CreatDelivery))
	deliveyRouter.Get("/_mget",hero.Handler(deliever.MgetDelivery))

}