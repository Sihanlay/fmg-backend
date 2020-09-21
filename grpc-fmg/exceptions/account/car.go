package accountException

import "grpc-demo/models"

func AccountCarNotFount() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5400,
		Message: "找不到该购物车",
	}
}
