package api

import (
	"gin_dev/app/model"
	"gin_dev/app/response"
	"github.com/gin-gonic/gin"
)

func GetArticleList(c *gin.Context){
	var article model.Article
	model.Db.First(&article)
	response.Success(c,"获取成功",article)
	return
}