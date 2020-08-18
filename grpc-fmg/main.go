package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"grpc-demo/config"
	"grpc-demo/core/cache"
	viewbase "grpc-demo/core/view"
	"grpc-demo/models/db"
	"grpc-demo/utils"
	"grpc-demo/utils/middlewares"
	"grpc-demo/views"
)

func initRouter(app *iris.Application) {
	views.RegisterAccountRouters(app)
}

func main() {
	app := iris.New()
	// 注册控制器
	app.UseGlobal(middlewares.AbnormalHandle, middlewares.RequestLogHandle)
	hero.Register(viewbase.ViewBase)
	// 注册路由
	initRouter(app)
	// 初始化配置
	config.InitConfig()
	utils.InitGlobal()
	// 初始化数据库
	db.InitDB()
	// 初始化缓存
	//cache.InitDijan()
	cache.InitRedisPool()
	// 初始化任务队列
	//queue.InitTaskQueue()
	// 启动系统
	app.Run(iris.Addr(":80"), iris.WithoutServerError(iris.ErrServerClosed))

}

//func main() {
//	// 初始化配置
//	config.InitConfig()
//	db.InitDB()
//	cache.InitRedisPool()
//
//	lis, err := net.Listen("tcp", ":6000")
//	if err != nil {
//		panic(err)
//	}
//    fmt.Println("hello")
//	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor.AbnormalHandle))
//	Address.RegisterAddressInServer(s, &views.AddressHandle{})
//	s.Serve(lis)
//}
