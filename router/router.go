package router

import (
	"gin_dev/controller/admin"
	"github.com/gin-gonic/gin"
)

func router(r *gin.Engine)*gin.Engine{
	//路由组
	article := r.Group("/article")
	{
		article.GET("/list",admin.Article{}.List)
		article.GET("/edit",admin.Article{}.Edit)
		article.GET("/create",admin.Article{}.Create)
		article.GET("/delete",admin.Article{}.Delete)
	}
	return r
}