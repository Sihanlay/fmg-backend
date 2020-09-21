package accountException

import "grpc-demo/models"

func AddressNotFount() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5303,
		Message: "找不到该地址",
	}
}
