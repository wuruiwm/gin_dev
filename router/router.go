package router

import "github.com/gin-gonic/gin"

func router(r *gin.Engine)*gin.Engine{
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}