package resource

import (
	"github.com/kataras/iris"
	"grpc-demo/core/auth"
	"grpc-demo/utils/qiniu"
)

// 获取七牛上传token
func GetQiNiuUploadToken(ctx iris.Context, auth authbase.AuthAuthorization) {
	// auth.CheckLogin()

	ctx.JSON(iris.Map{
		"token": qiniuUtils.GetUploadToken(),
	})
}
