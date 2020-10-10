package router

import (
	"gin_dev/config"
	"github.com/gin-gonic/gin"
)

func HttpInit(){
	r := gin.New()
	//初始化gin自身的两个全局中间件
	r.Use(gin.Logger(),gin.Recovery())
	//路由设置
	r = router(r)
	//启动服务
	_ = r.Run("0.0.0.0:"+config.GetString("server_port"))
}