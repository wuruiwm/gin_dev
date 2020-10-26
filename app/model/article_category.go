package model

type ArticleCategory struct {
	ID         uint   `gorm:"column:id;not null;primary_key;AUTO_INCREMENT;type:int(11)" json:"id"`
	Title      string `gorm:"column:title;not null;default:'';comment:'分类名称';type:char(255)" json:"title"`
	CreateTime int    `gorm:"column:create_time;not null;comment:'创建时间';type:int(11)" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time;not null;comment:'修改时间';type:int(11)" json:"update_time"`
}

func (*ArticleCategory) TableName() string {
	return `article_category`
}