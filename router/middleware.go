package router

import (
	"github.com/gin-gonic/gin"
)

func adminAuth()gin.HandlerFunc{
	return func(c *gin.Context){
		//前置操作
		//controller.Response{}.Return()
		c.Next()
		//后置操作
	}
}