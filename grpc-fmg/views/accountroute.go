package views

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"grpc-demo/views/account"
)

func RegisterAccountRouters(app *iris.Application) {

	// 题目路由
	placeTagRouter := app.Party("account/info")

	//placeTagRouter.Post("", hero.Handler(goods_tag.CreatePlaceTag))
	placeTagRouter.Get("/list", hero.Handler(account.GetAccountList))
	placeTagRouter.Put("/{uid:int}", hero.Handler(account.PutAccount))
	placeTagRouter.Post("/_mget", hero.Handler(account.MgetAccount))

	app.Party("account/").Post("/login/", hero.Handler(account.Login))
	//placeTagRouter.Put("/{pid:int}", hero.Handler(goods_tag.PutPlaceTag))
	//placeTagRouter.Delete("/{pid:int}", hero.Handler(goods_tag.DeletePlaceTag))
	//placeTagRouter.Post("/_mget", hero.Handler(goods_tag.MgetPlaceTag))
	//
	//saleTagRouter := app.Party("goods/sale_tag")
	//
	//saleTagRouter.Post("", hero.Handler(goods_tag.CreateSaleTag))
	//saleTagRouter.Get("/list", hero.Handler(goods_tag.ListSaleTag))
	//saleTagRouter.Put("/{sid:int}", hero.Handler(goods_tag.PutSaleTag))
	//saleTagRouter.Delete("/{sid:int}", hero.Handler(goods_tag.DeleteSaleTag))
	//saleTagRouter.Post("/_mget", hero.Handler(goods_tag.MgetSaleTag))

	//kindTagRouter := app.Party("goods/kind_tag")

	//kindTagRouter.Post("", hero.Handler(goods_tag.CreateKindTag))
	//kindTagRouter.Get("/list", hero.Handler(goods_tag.ListKindTag))
	//kindTagRouter.Put("/{kid:int}", hero.Handler(goods_tag.PutKindTag))
	//kindTagRouter.Delete("/{kid:int}", hero.Handler(goods_tag.DeleteKindTag))
	//kindTagRouter.Post("/_mget", hero.Handler(goods_tag.MgetKindTag))

}
