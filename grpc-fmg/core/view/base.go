package viewbase

import (
	authbase "grpc-demo/core/auth"
	"github.com/kataras/iris"
)

func ViewBase(ctx iris.Context) authbase.AuthAuthorization {
	return authbase.NewAuthAuthorization(&ctx)
}

