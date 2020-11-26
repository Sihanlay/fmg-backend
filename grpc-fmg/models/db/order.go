package db

import paramsUtils "grpc-demo/utils/params"

//总订单
type TestOrder struct {
	ID int `gorm:"primary_key" json:"id"`

	//拆分状态
	Status int16 `json:"status" gorm:"not null"`

	//用户id
	AccountID int `json:"account_id" gorm:"not null;index"`

	//地址id
	AddressID int `json:"address_id" gorm:"not null;index"`

	//微信支付关联id
	WxPayOrderId int `json:"wx_pay_order_id" gorm:"not null;index"`

	//付款时间
	PayTime int64 `json:"pay_time"`

	//是否付款
	PayOrNot bool `json:"pay_or_not"`

	//总优惠
	TotalCoupon int `json:"total_coupon" `

	//总运费
	TotalExpFare int `json:"total_exp_fare"`

	//商品总额
	TotalGoodsAmount int `json:"total_goods_amount" gorm:"not null"`

	//实付订单总额
	TotalOrderAmount int `json:"total_order_amount" gorm:"not null"`

	//订单号
	OrderNum string `json:"order_num" gorm:"not null"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

//子订单
type TestChildOrder struct {
	ID int `gorm:"primary_key" json:"id"`

	//状态
	OrderStatus int16 `json:"order_status" gorm:"not null"`

	//订单号
	OrderNum string `json:"order_num" gorm:"not null"`

	//用户id
	AccountID int `json:"account_id" gorm:"not null;index"`

	//地址id
	AddressID int `json:"address_id" gorm:"not null;index"`

	//总订单id
	OrderID int `json:"order_id" gorm:"not null;index"`

	//发货方式
	Delivery int16 `json:"delivery" gorm:"not null"`

	//子订单总优惠
	ChildTotalCoupon int `json:"child_total_coupon" `

	//子订单商品总额
	ChildGoodsAmount int `json:"child_goods_amount" gorm:"not null"`

	//子订单运费
	ChildExpFare int `json:"child_exp_fare"`

	//子订单实付金额
	ChildOrderAmount int `json:"child_order_amount" gorm:"not null"`

	//发货时间
	DeliveryTime int64 `json:"delivery_time"`

	//收货时间
	GetTime int64 `json:"get_time"`

	//快递单号
	TrackingID string `json:"tracking_id"`
	
	//快递公司
	TrackingCompany string `json:"tracking_company"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

//订单明细
type TestOrderDetail struct {
	ID int `gorm:"primary_key" json:"id"`

	//总订单id
	OrderID int `json:"order_id" gorm:"not null;index"`

	//子订单id
	ChildOrderID int `json:"child_order_id" gorm:"not null;index"`

	//商品id
	GoodsID int `json:"goods_id" gorm:"not null;index"`

	//商品规格
	GoodsSpecificationID int `json:"goods_specification_id" gorm:"not null"`

	//商品数量
	PurchaseQty int `json:"purchase_qty" gorm:"not null"`

	//备注
	Message string `json:"message"`

	//发货方式
	//Delivery int16 `json:"delivery" gorm:"not null"`

	//优惠
	Coupon int `json:"coupon" `

	//运费
	ExpFare int `json:"exp_fare"`

	//商品总额
	GoodsAmount int `json:"goods_amount" gorm:"not null"`

	//实付订单总额
	//OrderAmount float32 `json:"order_amount" gorm:"not null"`
	
	//是否评价
	IsComment int `json:"is_comment"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

//子订单备份
type TestChildOrderCopy struct {
	ID int `gorm:"primary_key" json:"id"`

	//状态
	OrderStatus int16 `json:"order_status" gorm:"not null"`

	//订单号
	OrderNum string `json:"order_num" gorm:"not null"`

	//用户id
	AccountID int `json:"account_id" gorm:"not null;index"`

	//地址id
	AddressID int `json:"address_id" gorm:"not null;index"`

	//总订单id
	OrderID int `json:"order_id" gorm:"not null;index"`

	//发货方式
	Delivery int16 `json:"delivery" gorm:"not null"`

	//子订单总优惠
	ChildTotalCoupon int `json:"child_total_coupon" `

	//子订单商品总额
	ChildGoodsAmount int `json:"child_goods_amount" gorm:"not null"`

	//子订单运费
	ChildExpFare int `json:"child_exp_fare"`

	//子订单实付金额
	ChildOrderAmount int `json:"child_order_amount" gorm:"not null"`

	//发货时间
	DeliveryTime int64 `json:"delivery_time"`

	//收货时间
	GetTime int64 `json:"get_time"`

	//快递单号
	TrackingID string `json:"tracking_id"`

	//快递公司
	TrackingCompany string `json:"tracking_company"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

////订单备份
////type TestOrderCopy struct {
////	ID int `gorm:"primary_key" json:"id"`
////
////	//订单号
////	OrderNum string `json:"order_num" gorm:"not null"`
////
////	//用户id
////	AccountID int `json:"account_id" gorm:"not null;index"`
////
////	//地址id
////	AddressID int `json:"address_id" gorm:"not null;index"`
////
////	//状态
////	Status int16 `json:"status" gorm:"not null"`
////
////	//微信支付关联id
////	WxPayOrderId int `json:"wx_pay_order_id" gorm:"not null;index"`
////
////	//付款时间
////	PayTime int64 `json:"pay_time"`
////
////	//总优惠
////	TotalCoupon float32 `json:"total_coupon" `
////
////	//总运费
////	TotalExpFare float32 `json:"total_exp_fare"`
////
////	//商品总额
////	TotalGoodsAmount float32 `json:"total_goods_amount" gorm:"not null"`
////
////	//实付订单总额
////	TotalOrderAmount float32 `json:"total_order_amount" gorm:"not null"`
////
////	// 创建时间
////	CreateTime int64 `json:"create_time"`
////
////	// 更新时间
////	UpdateTime int64 `json:"update_time"`
////}

var orderField = []string{
	"ID", "AccountID", "AddressID", "WxPayOrderId", "PayTime", "TotalCoupon", "TotalExpFare",
	"TotalGoodsAmount", "TotalOrderAmount", "CreateTime", "UpdateTime", "Status", "OrderNum",
}

var childOrderField = []string{
	"ID", "OrderID", "Delivery", "ChildGoodsAmount", "ChildTotalCoupon","ChildExpFare","ChildOrderAmount", "CreateTime", "UpdateTime", "TrackingID", "DeliveryTime", "GetTime",
	"AccountID", "AddressID", "OrderStatus", "OrderNum","TrackingCompany",
}

var orderDetailField = []string{
	"ID", "ChildOrderID", "GoodsID", "GoodsSpecificationID", "PurchaseQty", "Message", "OrderID",
	"Coupon", "ExpFare", "GoodsAmount", "CreateTime", "UpdateTime",
}

func (o TestChildOrder) GetInfo() map[string]interface{} {

	v := paramsUtils.ModelToDict(o, childOrderField)

	//var address Address


	var order TestOrder
	//Driver.GetOne("test_order", o.OrderID, &order)
	Driver.Where("id = ?",o.OrderID).First(&order)
	v["test_order"] = paramsUtils.ModelToDict(order, orderField)

	var orderDetails []TestOrderDetail
	if err := Driver.Where("child_order_id = ?", o.ID).Find(&orderDetails).Error; err == nil {
		orderDetailData := make([]map[string]interface{}, len(orderDetails))
		for index, detail := range orderDetails {
			v1 := paramsUtils.ModelToDict(detail, orderDetailField)
			orderDetailData[index] = v1
		}
		v["order_detail"] = orderDetailData
	}
	return v
}


func (o TestOrder) GetInfo() map[string]interface{}{
	v := paramsUtils.ModelToDict(o, orderField)
	//"child_order":[{"id":,"order_id":,"order_detail":[]},{}]

	var childOrders []TestChildOrder
	if err := Driver.Where("order_id = ?", o.ID).Find(&childOrders).Error; err == nil {
		childOrderData := make([]map[string]interface{}, len(childOrders))
		for i, c := range childOrders {
			v1 := paramsUtils.ModelToDict(c, childOrderField)
			var orderDetails []TestOrderDetail
			if err := Driver.Where("child_order_id = ?", c.ID).Find(&orderDetails).Error; err == nil {
				orderDetailData := make([]map[string]interface{}, len(orderDetails))
				for index, detail := range orderDetails {
					v2 := paramsUtils.ModelToDict(detail, orderDetailField)
					orderDetailData[index] = v2
				}
				v1["order_detail"] = orderDetailData
			}
			childOrderData[i] = v1
		}
		v["child_order"] = childOrderData
	}

	return v
}
