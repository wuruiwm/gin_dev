package model

import (
	"encoding/json"
	"gin_dev/common"
	"time"
)

type Article struct {
	ID         uint   `gorm:"column:id;not null;primary_key;AUTO_INCREMENT;type:int(11)" json:"id"`
	CategoryID int    `gorm:"column:category_id;not null;comment:'分类id';type:int(11)" json:"category_id"`
	Title      string `gorm:"column:title;not null;default:'';comment:'标题';type:char(255)" json:"title"`
	Content    string `gorm:"column:content;not null;comment:'内容';type:text" json:"content"`
	CreateTime int    `gorm:"column:create_time;not null;comment:'创建时间';type:int(11)" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time;not null;comment:'修改时间';type:int(11)" json:"update_time"`
}

type ArticleListResult struct {
	Article
	CategoryTitle      string `gorm:"column:category_title;not null;default:'';comment:'分类名称';type:char(255)" json:"category_title"`
}

func (*Article) TableName() string {
	return `article`
}

func ArticleList(page int,limit int,where map[string]interface{},keyword string)([]*ArticleListResult, error){
	offset := (page - 1)*limit
	var articleListResult []*ArticleListResult
	articleModel := db.Table("article as a").
		Joins("LEFT JOIN article_category ac on a.category_id=ac.id").
		Where(where)
	if keyword != ""{
		articleModel = articleModel.Where("a.title like ?","%"+keyword+"%")
	}
	err := articleModel.Where("a.title like ?","%"+keyword+"%").
		Order("a.id desc").
		Offset(offset).
		Limit(limit).
		Select("a.id,a.category_id,a.title,ac.title as `category_title`,a.create_time,a.update_time").
		Find(&articleListResult).Error
	return articleListResult,err
}

func ArticleCreate(title string,content string)error{
	article := Article{
		Title: title,
		Content: content,
		CreateTime: common.GetUnixTime(),
		UpdateTime: common.GetUnixTime(),
	}
	return db.Create(&article).Error
}

func ArticleUpdate(id int,title string,content string)error{
	if _,err := ArticleDetail(id);err != nil{
		return err
	}
	article := Article{
		Title: title,
		Content: content,
		UpdateTime: int(time.Now().Unix()),
	}
	return db.Model(&Article{}).
		Where("id",id).
		Updates(&article).Error
}

func ArticleDelete(id int)error{
	if _,err := ArticleDetail(id);err != nil{
		return err
	}
	return db.Where("id",id).
		Delete(&Article{}).
		Error
}

func ArticleDetail(id int)(*Article,error){
	var (
		article Article
		err error
	)
	err = db.Take(&article,id).Error
	return &article,err
}

func ArticleCategoryAll()([]*ArticleCategory,error){
	var (
		articleCategory []*ArticleCategory
		err error
	)
	//判断是否有缓存 有缓存则直接返回
	cacheKey := "cache:article_category_all"
	articleCategoryJson := common.GetCache(cacheKey)
	if articleCategoryJson != ""{
		if err := json.Unmarshal([]byte(articleCategoryJson),&articleCategory);err == nil{
			return articleCategory,nil
		}
	}
	//无缓存 查询结果缓存后返回
	if err = db.Select("id,title").Find(&articleCategory).Error;err != nil{
		return articleCategory,err
	}
	if buf,err := json.Marshal(articleCategory);err == nil{
		_ = common.SetCache(cacheKey,string(buf),7200)
	}
	return articleCategory,nil
}