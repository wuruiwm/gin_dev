package model

type Article struct {
	ID         uint   `gorm:"column:id;not null;primary_key;AUTO_INCREMENT;type:int(11)" json:"id"`
	CategoryID int    `gorm:"column:category_id;not null;default:'0';comment:'分类id';type:int(11)" json:"category_id"`
	Title      string `gorm:"column:title;not null;default:'';comment:'标题';type:char(255)" json:"title"`
	Content    string `gorm:"column:content;not null;comment:'内容';type:text" json:"content"`
	CreateTime int    `gorm:"column:create_time;not null;default:'0';comment:'创建时间';type:int(11)" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time;not null;default:'0';comment:'修改时间';type:int(11)" json:"update_time"`
}


func (*Article) TableName() string {
	return `article`
}