package db

//资讯
type News struct{
	ID int `gorm:"primary_key" json:"id"`

	//标题
	Title string `json:"title" gorm:"not null"`

	//内容
	Content string `json:"content" gorm:"type: text"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

//todo 资讯标签
type NewsTag struct{
	ID int `gorm:"primary_key" json:"id"`

	//名称
	Name string `json:"name" gorm:"not null`
}

//标签挂载
type NewsAndTag struct{
	ID int `gorm:"primary_key" json:"id"`

	//资讯id
	NewsID int `json:"news_id" gorm:"not null;index"`

	//标签id
	TagID int `json:"tag_id" gorm:"not null;index"`
}