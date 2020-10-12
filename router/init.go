package router

import (
	"fmt"
	"gin_dev/config"
	"github.com/gin-gonic/gin"
)

//初始化gin的http服务
func HttpInit(){
	//是否开启debug
	if config.GetBool("debug") {
		gin.SetMode(gin.DebugMode)
	}else{
		gin.SetMode(gin.ReleaseMode)
	}
	//实例化gin
	r := gin.Default()
	//路由设置
	r = router(r)
	//启动服务
	_ = r.Run(getHttpString())
}

func getHttpString()string{
	port := config.GetString("server_port")
	return fmt.Sprintf("0.0.0.0:%s",port)
}