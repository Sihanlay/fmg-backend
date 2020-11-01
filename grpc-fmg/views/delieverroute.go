package views

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"grpc-demo/views/deliever"
)

func DeliveryRouters(app *iris.Application) {

	deliveyRouter := app.Party("delivery/info")
	deliveyRouter.Post("", hero.Handler(deliever.CreatDelivery))
	deliveyRouter.Post("/post", hero.Handler(deliever.DeliveryInfo))
	deliveyRouter.Post("/_mget", hero.Handler(deliever.MgetDelivery))
	deliveyRouter.Get("/list", hero.Handler(deliever.GetDeliveryList))
}
