package accountException

import (
	"grpc-demo/models"
)

func AuthIsNotLogin() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5300,
		Message: "尚未登录",
	}
}

func NoPermission() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5301,
		Message: "无权限执行此操作",
	}
}

func LoginFailed() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5302,
		Message: "未注册到数据库",
	}
}

func AccountNotFount() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5303,
		Message: "找不到该用户",
	}
}

func NotClientModel() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5304,
		Message: "不是client模式",
	}
}
