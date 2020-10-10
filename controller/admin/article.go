package admin

import (
	"gin_dev/controller"
	"github.com/gin-gonic/gin"
)

type Article struct {

}

func (article Article)List(c *gin.Context){
	controller.Success(c,"获取成功",nil,nil)
}