package router

import (
	"gin_dev/app/controller/api"
	"gin_dev/app/middleware"
	"github.com/gin-gonic/gin"
)

//设置路由
func router(r *gin.Engine)*gin.Engine{
	//路由组
	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.UserAuth())
	{
		article := adminGroup.Group("/article")
		{
			article.GET("/list", api.ArticleList)
			article.POST("/create", api.ArticleCreate)
			article.POST("/update", api.ArticleUpdate)
			article.POST("/delete", api.ArticleDelete)
		}
	}
	return r
}