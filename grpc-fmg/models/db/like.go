package db

type Like struct {

	Id int `gorm:"primary_key" json:"id"`

	AccountId int `json:"account_id" `

	GoodsId int `json:"goods_id"`

	GoodsSpecification string `json:"goods_specification"`

	GoodsSpecificationId int `json:"goods_specification_id"`

	CreateTime int64 `json:"create_time"`

}
