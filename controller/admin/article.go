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
func (article Article)Edit(c *gin.Context){
	article.Success(c,"获取成功",nil)
}
func (article Article)Create(c *gin.Context){
	article.Success(c,"获取成功",nil)
}
func (article Article)Delete(c *gin.Context){
	article.Success(c,"获取成功",nil)
}