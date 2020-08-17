package db

type Goods struct{
	ID int `gorm:"primary_key" json:"id"`

	//商品名
	Name string `json:"name" gorm:"not null"`

	//商品卖点
	SalePoint string `json:"sale_point" gorm:"default: null"`

	//发货地
	Address string `json:"address" gorm:"default: null"`

	//月销量
	MouthSale int `json:"mouth_sale" `

	//付款人数
	People int `json:"people" `

	//显示价格
	Price float64 `json:"price" gorm:"default: null"`

	//是否有优惠
	Sale bool `json:"sale" gorm:"default: false"`

	//原价
	PrivatePrice float64 `json:"private_price" gorm:"not null"`

	//库存扣减方式
	TakeRemove bool `json:"take_remove" gorm:"default: false"`

	//库存
	Total int `json:"total" `

	//是否展示库存
	ShowTotal bool `json:"show_total" gorm:"default: false"`

	//运费
	Carriage float64 `json:"carriage" gorm:"default: null"`

	//收货方式
	GetWay int16 `json:"get_way" gorm:"not null"`


	//上架时间
	PutawayTime int64 `json:"putaway_time" gorm:"default: null"`

	//上架方式
	Putaway int16 `json:"putaway" gorm:"not null"`

	//商品状态
	OnSale bool `json:"on_sale" gorm:"default: false"`

	//是否支持换货
	Exchange bool `json:"exchange" gorm:"default: false"`

	//是否支持7天无理由退货
	SaleReturn bool `json:"sale_return" gorm:"default: false"`

	//起售数量
	MinSale int `json:"min_sale" gorm:"default: 1"`

	//是否预售
	Advance bool `json:"advance" gorm:"default: false"`

	//预售时间
	AdvanceTime int64 `json:"advance_time" gorm:"default: null"`

	//是否限购
	Limit bool `json:"limit" gorm:"default: false"`

	//限购数量
	LimitTotal int `json:"limit_total" gorm:"default: null"`

	//模版id
	TemplateID int `json:"template_id" gorm:"not null"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}
