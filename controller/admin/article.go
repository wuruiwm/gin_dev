package admin

import (
	"github.com/gin-gonic/gin"
)

type Article struct {

}

func (article Article)List(c *gin.Context){
	Response.Success(Response{},c,"dsada",nil)
}
func (article Article)Edit(c *gin.Context){
	Response.Success(c,"获取成功",nil)
}
func (article Article)Create(c *gin.Context){
	Response.Success(c,"获取成功",nil)
}
func (article Article)Delete(c *gin.Context){
	Response.Success(c,"获取成功",nil)
}