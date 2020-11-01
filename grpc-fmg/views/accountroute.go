package views

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"grpc-demo/views/account"
	"grpc-demo/views/address"
	"grpc-demo/views/comment"
)

func RegisterAccountRouters(app *iris.Application) {

	// 账户信息路由
	accountRouter := app.Party("account/info")

	accountRouter.Get("/list", hero.Handler(account.GetAccountList))
	accountRouter.Put("/{uid:int}", hero.Handler(account.PutAccount))
	accountRouter.Post("/_mget", hero.Handler(account.MgetAccount))

	//登录路由
	loginRouter := app.Party("account/login")
	loginRouter.Post("/wx_login", hero.Handler(account.Login))
	loginRouter.Post("/register", hero.Handler(account.Register))
	loginRouter.Post("/web_login", hero.Handler(account.WebLogin))
	loginRouter.Post("/logout", hero.Handler(account.Logout))

	//地址路由
	addressRouter := app.Party("address/info")
	addressRouter.Post("/{uid:int}", hero.Handler(address.CreatAddress))
	addressRouter.Get("/{uid:int}", hero.Handler(address.MGetAddress))
	addressRouter.Put("/{aid:int}", hero.Handler(address.PutAddress))
	addressRouter.Get("/list", hero.Handler(address.ListAddress))
	addressRouter.Get("/get/{aid:int}", hero.Handler(address.GetAddress))
	addressRouter.Delete("/delete/{aid:int}", hero.Handler(address.DeleteAddress))

	//购物车路由
	accountCarRouter := app.Party("car/info")
	accountCarRouter.Put("/_mset/{uid:int}", hero.Handler(account.MsetAccountCar))
	accountCarRouter.Post("/{uid:int}/{gid:int}", hero.Handler(account.CreatAccountCar))
	accountCarRouter.Post("/_mget/{uid:int}", hero.Handler(account.MgetAccountCar))
	accountCarRouter.Delete("/delete/{cid:int}", hero.Handler(account.DeleteGoodsCar))
	accountCarRouter.Put("/put/{cid:int}", hero.Handler(account.PutAccountCar))

	//评价路由
	commentRouter := app.Party("comment/info")
	commentRouter.Post("/{uid:int}", hero.Handler(comment.CreatComment))
	commentRouter.Get("/get/{cid:int}", hero.Handler(comment.GetComment))
	commentRouter.Put("/put/{cid:int}", hero.Handler(comment.PutComment))
	commentRouter.Delete("/delete/{cid:int}", hero.Handler(comment.DeleteComment))
	commentRouter.Get("/_mget/{cid:int}", hero.Handler(comment.MgetComment))
}
