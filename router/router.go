package router

import (
	"github.com/gin-gonic/gin"
)

//设置路由
func router(r *gin.Engine)*gin.Engine{
	r = adminRouter(r)
	return r
}