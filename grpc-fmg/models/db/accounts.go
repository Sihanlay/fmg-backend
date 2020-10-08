package db

// 用户主账户表
type Account struct {
	Id int `gorm:"primary_key" json:"id"`

	// 昵称
	Nickname string `json:"nickname"`

	// 角色
	Role int16 `json:"role" gorm:"not null"`

	// 电话
	Phone string `json:"phone"`

	// 电话验证与否
	PhoneValidated bool `json:"phone_validated" gorm:"default:false"`

	//微信号
	WXid string `json:"wxid" gorm:"not null"`

	////积分
	//Usrscores Scores `json:"-" gorm:"ForeignKey:AccountId"`

	////卡券
	//Usrcard []Cards `json:"-" gorm:"ForeignKey:AccountId"`

	//生日
	Birthday int64 `json:"birthday"`

	//// 通过第三方进来的，首次设定密码不需要给旧密码
	//Init bool `json:"init" gorm:"default:false"`

	// 邮箱
	Email string `json:"email"`

	// 邮箱验证与否
	EmailValidated bool `json:"email_validated" gorm:"default:false"`

	// 头像
	Avator string `json:"avator"`

	// 一句话签名
	Motto string `json:"motto"`

	//省
	Province string `json:"province"`

	//市
	City string `json:"city"`

	//国
	Country string `json:"country"`

	// 设置 保留字段
	Options string `json:"options" gorm:"default:''"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`

	OpenId string `json:"open_id"`

	SessionKey string `json:"session_key"`

	// ----------------------------------------
}

// 用户第三方验证
//type AccountOauth struct {
//
//	Id        int `gorm:"primary_key" json:"id"`
//
//	// 用户id
//	AccountId int `json:"account_id"`
//
//	// 关联模块
//	Model int16 `json:"model"`
//
//	// openid
//	OpenId string `json:"open_id"`
//
//	// 用户信息
//	UserInfo string `json:"user_info" gorm:"default:'';type:text"`
//
//	// 其他信息
//	ExtraInfo string `json:"extra_info" gorm:"default:'';type:text"`
//
//	// 创建时间
//	CreateTime int64 `json:"create_time"`
//
//	// 更新时间
//	UpdateTime int64 `json:"update_time"`
//}

//用户积分表
type Scores struct {
	Id int `gorm:"primary_key" json:"id"`

	//积分
	Scores int64 `json:"scores"`

	// 用户id
	AccountId int `json:"account_id"`

	//积分来源 1：订单 2：评价 3：订单取消返还 4：拒收返还
	ScoresSrc int16 `json:"scoresSrc"`

	//积分来源订单
	OrderId int `json:"orderid"`

	//积分类型（收入支出）
	ScoreType int16 `json:"scoretype"`

	//描述
	ScoreRemark string `json:"scoreremark" gorm:"default:'';type:text"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

//用户优惠券表
type Cards struct {
	Id int `gorm:"primary_key" json:"id"`

	// 用户id
	AccountId int `json:"account_id"`

	// 卡券号
	CardNo string `json:"cardno"`

	// 订单ID
	OrderId int `json:"orderid"`

	//是否使用
	UseValid bool `json:"usedvalid" gorm:"default:true"`

	//是否有效 有效标志 1：有效 -1：过期
	DateFlag bool `json:"date_flag"`

	// 创建时间
	CreateTime int64 `json:"create_time"`
}

//用户购物车表
type AccountCar struct {
	Id int `gorm:"primary_key" json:"id"`

	AccountId int     `json:"account_id" `

	GoodsId int   `json:"goods_id"`

	GoodsCount int `json:"goods_count" gorm:"default:0"`

	GoodsSpecification string `json:"goods_specification"`

	GoodsSpecificationId int `json:"goods_specification_id"`

	//GoodsName string `json:"goods_name"`
	//
	//GoodsPrice int `json:"goods_price"`

	IsCheck bool `json:"is_check" gorm:"default:true"`

	//Picture string `json:"picture"`

	CreateTime int64 `json:"create_time"`
}
