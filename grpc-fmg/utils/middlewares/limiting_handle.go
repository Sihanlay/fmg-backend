package middlewares

import (
	"grpc-demo/utils/current_limiting"
	"github.com/kataras/iris"
)

// 限流
func LimitingHandle(ctx iris.Context) {
	// 获取令牌
	<-current_limiting.TokenBucket
	ctx.Next()
}
