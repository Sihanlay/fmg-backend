package systemException

import (
	"grpc-demo/models"
)

func SystemException() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status: false,
		ErrCode: 5000,
		Message: "系统错误",
	}
}

func sqlException() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status: false,
		ErrCode: 5001,
		Message: "数据库操作错误, 所有操作已回滚",
	}
}


