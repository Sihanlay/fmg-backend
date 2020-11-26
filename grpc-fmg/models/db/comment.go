package db

//评价
type Comment struct {
	ID int `gorm:"primary_key" json:"id"`

	// 发布人
	AuthorID int `json:"author_id" gorm:"not null;index"`

	//对应商品
	GoodID int `json:"good_id" gorm:"not null;index"`

	//评价标签 1好评 2中评 3差评
	CommentTag int `json:"comment_tag"`
	//评价内容
	Content string `json:"content"`
	//图片路径
	Pictures string `json:"pictures"`

	//是否有图
	//HasPicture bool `json:"has_picture" gorm:"default: false"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

//评价标签
type CommentTag struct {
	ID int `gorm:"primary_key" json:"id"`

	// 标题
	Title string `json:"title" gorm:"not null"`
}

//评价挂载标签（多对多）
type CommentAndTag struct {
	ID int `gorm:"primary_key" json:"id"`

	//评价id
	CommentID int `json:"comment_id" gorm:"not null;index"`

	//标签id
	CommentTagID int `json:"comment_tag_id" gorm:"not null;index"`
}
