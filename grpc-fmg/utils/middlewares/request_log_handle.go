package middlewares

import (
	"grpc-demo/utils/log"
	"github.com/kataras/iris"
)

// 请求日志
func RequestLogHandle(ctx iris.Context) {
	logUtils.Println("Host:", ctx.RemoteAddr(), "Method:", ctx.Method(), "Path:", ctx.Path())
	ctx.Next()
}
