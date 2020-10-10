package router

import (
	"gin_dev/controller/admin"
	"github.com/gin-gonic/gin"
)

//设置路由
func router(r *gin.Engine)*gin.Engine{
	//路由组
	adminGroup := r.Group("/admin")
	adminGroup.Use(adminAuth())
	{
		article := adminGroup.Group("/article")
		{
			article.GET("/list", admin.Article{}.List)
			article.GET("/edit", admin.Article{}.Edit)
			article.GET("/create", admin.Article{}.Create)
			article.GET("/delete", admin.Article{}.Delete)
		}
	}
	return r
}