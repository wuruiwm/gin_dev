package router

import (
	"gin_dev/config"
	"github.com/gin-gonic/gin"
)

func init(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run("0.0.0.0:"+config.GetString("server_port"))
}