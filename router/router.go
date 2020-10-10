package router

import (
	"gin_dev/controller/admin"
	"github.com/gin-gonic/gin"
)

func router(r *gin.Engine)*gin.Engine{
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":0,
			"msg":"hello world",
		})
	})
	r.GET("/article/list",admin.Article{}.List)
	return r
}