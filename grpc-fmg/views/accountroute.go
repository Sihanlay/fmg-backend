package views

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"grpc-demo/views/account"
	"grpc-demo/views/address"
)

func RegisterAccountRouters(app *iris.Application) {

	// 账户信息路由
	accountRouter := app.Party("account/info")

	//placeTagRouter.Post("", hero.Handler(goods_tag.CreatePlaceTag))
	accountRouter.Get("/list", hero.Handler(account.GetAccountList))
	accountRouter.Put("/{uid:int}", hero.Handler(account.PutAccount))
	accountRouter.Post("/_mget", hero.Handler(account.MgetAccount))

	app.Party("account/").Post("/login/", hero.Handler(account.Login))

	//地址路由
	addressRouter := app.Party("address/info")
	addressRouter.Post("/{uid:int}", hero.Handler(address.CreatAddress))
	addressRouter.Get("/{uid:int}", hero.Handler(address.GetAddress))
	addressRouter.Put("/{aid:int}", hero.Handler(address.PutAddress))

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
