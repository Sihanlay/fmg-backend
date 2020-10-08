package db

//发货信息

type Delivery struct {

	ID int `gorm:"primary_key" json:"id"`

	OrderCode int `json:"order_code"`

	DeliveryCorpName string `json:"delivery_corp_name"`

	DeliverySheetCode int `json:"delivery_sheet_code"`

	InvoiceStatus int `json:"invoice_status"`

	TaskId int `json:"task_id"`

	ReceiveCode string `json:"receive_code"`

}
