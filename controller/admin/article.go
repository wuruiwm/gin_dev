package admin

import (
	"gin_dev/controller"
	"github.com/gin-gonic/gin"
)

type Article struct {
	*controller.Response
}

func (article Article)List(c *gin.Context){
	article.Success(c,"获取成功",nil)
}