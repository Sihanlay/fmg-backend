package accountException

import (
	"grpc-demo/models"
)

func AuthIsNotLogin() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status: false,
		ErrCode: 5300,
		Message: "尚未登录",
	}
}

func NoPermission() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status: false,
		ErrCode: 5301,
		Message: "无权限执行此操作",
	}
}

func AccountNotFount() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status: false,
		ErrCode: 5302,
		Message: "找不到该用户",
	}
}




