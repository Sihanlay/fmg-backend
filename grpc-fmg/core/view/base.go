package viewbase

import (
	"github.com/kataras/iris"
	authbase "grpc-demo/core/auth"
)

func ViewBase(ctx iris.Context) authbase.AuthAuthorization {
	return authbase.NewAuthAuthorization(ctx)
}
