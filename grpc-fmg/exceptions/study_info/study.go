package study_info

import "grpc-demo/models"

func InfoNotExists() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5302,
		Message: "资讯不存在",
	}
}
