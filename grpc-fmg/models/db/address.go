package db

//地址
type Address struct {
	ID        int     `gorm:"primary_key" json:"id"`

	AccountId int     `json:"account_id"`
	//详细地址
	Detail string `json:"name" gorm:"not null;type:text"`

	//国

	CountryId int     `json:"country_id"`
	//省

	ProvinceID int      `json:"province_id"`
	//城市id

	CityID int  `json:"city_id"`


	DistrictID int      `json:"District"`

	Phone string `json:"phone"`

	Name string `json:"name"`
	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

//城市
type Country struct {
	ID int `gorm:"primary_key" json:"id"`

	// 标题
	Name string `json:"Countryname" gorm:"not null"`

	Code int `json:"country_code" gorm:"not null;index"`
}

type Province struct {
	ID int `gorm:"primary_key" json:"id"`

	//国家id
	CountryID int `json:"Country_id" gorm:"not null;index"`

	//省编码
	Code int `json:"Province_code" gorm:"not null;index"`

	Name string `json:"Provincename"`
}

type City struct {
	ID int `gorm:"primary_key" json:"id"`

	CountryID int `json:"Province_id" gorm:"not null;index"`

	//省id
	ProvinceID int `json:"Province_id" gorm:"not null;index"`
	//城市编码
	Code int `json:"City_code" gorm:"not null;index"`

	Name string `json:"Cityname"`
}

type District struct {
	ID int `gorm:"primary_key" json:"id"`

	CountryID int `json:"Province_id" gorm:"not null;index"`

	//省id
	ProvinceID int `json:"Province_id" gorm:"not null;index"`
	//城市id
	CityID int `json:"City_id" gorm:"not null;index"`
	//城市编码
	Code int `json:"City_code" gorm:"not null;index"`

	Name string `json:"Cityname"`
}
