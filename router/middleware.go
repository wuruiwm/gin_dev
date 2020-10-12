package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func adminAuth()gin.HandlerFunc{
	return func(c *gin.Context){
		//前置操作
		if c.Query("a") == "" {
			c.Abort() // 终止调用链条
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "a参数有问题，请检查参数",
			})
			return
		}
		c.Next()
		//后置操作
	}
}