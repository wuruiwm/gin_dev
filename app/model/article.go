package model

import "time"

type Article struct {
	ID         uint   `gorm:"column:id;not null;primary_key;AUTO_INCREMENT;type:int(11)" json:"id"`
	CategoryID int    `gorm:"column:category_id;not null;comment:'分类id';type:int(11)" json:"category_id"`
	Title      string `gorm:"column:title;not null;default:'';comment:'标题';type:char(255)" json:"title"`
	Content    string `gorm:"column:content;not null;comment:'内容';type:text" json:"content"`
	CreateTime int    `gorm:"column:create_time;not null;comment:'创建时间';type:int(11)" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time;not null;comment:'修改时间';type:int(11)" json:"update_time"`
}

func (*Article) TableName() string {
	return `article`
}

func ArticleList(page int,limit int,where map[string]interface{},keyword string)([]*Article, error){
	var article []*Article
	offset := (page - 1)*limit
	err := db.Where(where).Where("title like ?","%"+keyword+"%").Offset(offset).Limit(limit).Find(&article).Error
	if err == nil{
		return article,nil
	}
	return article,err
}

func ArticleCreate(title string,content string)error{
	article := Article{
		Title: title,
		Content: content,
		CreateTime: time.Now().Second(),
		UpdateTime: time.Now().Second(),
	}
	return db.Create(&article).Error
}