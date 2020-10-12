package api

import (
	"gin_dev/app/response"
	"github.com/gin-gonic/gin"
)

func GetArticleList(c *gin.Context){
	response.Success(c,"获取成功",nil)
	return
}