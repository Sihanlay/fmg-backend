package accountEnums

import enumsbase "grpc-demo/enums"

const (
	GetWayExpressage = 1  // 快递
	GetWaySameCity   = 2  // 同城配送
	GetWaySelf = 4 //自提


)

func NewGetWayEnums() enumsbase.EnumBaseInterface {
	return enumsbase.EnumBase{
		Enums: []int{1, 2, 4},
	}
}
