package middleware

import (
	"github.com/gin-gonic/gin"
)

func UserAuth()gin.HandlerFunc{
	return func(c *gin.Context){
		//前置操作
		//if(true){
		//	c.Abort()
		//	response.Return(c,response.NotLoginCode,"未登录",nil)
		//	return
		//}
		c.Next()
		//后置操作
	}
}