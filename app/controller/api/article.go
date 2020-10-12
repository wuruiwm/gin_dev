package api

import (
	"fmt"
	"gin_dev/app/model"
	"gin_dev/app/response"
	"github.com/gin-gonic/gin"
)

func GetArticleList(c *gin.Context){
	var article model.Article
	model.Db.First(&article)
	fmt.Println(article)
	response.Success(c,"获取成功",nil)
	return
}