package router

import (
	"gin_dev/app/controller/admin"
	"gin_dev/app/middleware"
	"github.com/gin-gonic/gin"
)

//设置路由
func router(r *gin.Engine)*gin.Engine{
	//路由组
	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.AdminAuth())
	{
		article := adminGroup.Group("/article")
		{
			article.GET("/", admin.ArticleList)
			article.POST("/", admin.ArticleCreate)
			article.PUT("/:id", admin.ArticleUpdate)
			article.DELETE("/:id", admin.ArticleDelete)
			article.GET("/:id", admin.ArticleDetail)
			adminGroup.GET("/articleCategoryAll", admin.ArticleCategory)
		}
	}
	return r
}