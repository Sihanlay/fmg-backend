package delievery
import "grpc-demo/models"

func OrderNotExists() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5302,
		Message: "子订单不存在",
	}
}

func SaleTagIsNotExists() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5302,
		Message: "销售标签不存在",
	}
}