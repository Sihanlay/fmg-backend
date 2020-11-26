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
	//accountCarRouter.Put("/_mset", hero.Handler(account.MsetAccountCar))
	accountCarRouter.Post("/{gid:int}", hero.Handler(account.CreatAccountCar))
	accountCarRouter.Post("/_mget", hero.Handler(account.MgetAccountCar))
	accountCarRouter.Delete("/delete", hero.Handler(account.MDeleteGoodsCar))
	accountCarRouter.Put("/put/{cid:int}", hero.Handler(account.PutAccountCar))

	//评价路由
	commentRouter := app.Party("comment/info")
	commentRouter.Post("/{gid:int}/{oid:int}", hero.Handler(comment.CreatComment))
	commentRouter.Post("/get/{gid:int}", hero.Handler(comment.MGetCommentByGood))
	commentRouter.Put("/put/{cid:int}", hero.Handler(comment.PutComment))
	commentRouter.Delete("/delete/{cid:int}", hero.Handler(comment.DeleteComment))
	commentRouter.Post("/_mget", hero.Handler(comment.MgetComment))

}
