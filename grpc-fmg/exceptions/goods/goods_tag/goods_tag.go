package goodsTagException

import "grpc-demo/models"

func PlaceTagIsNotExists() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5302,
		Message: "属地标签不存在",
	}
}

func SaleTagIsNotExists() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5302,
		Message: "销售标签不存在",
	}
}