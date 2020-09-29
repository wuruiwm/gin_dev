package router

import (
	"gin_dev/config"
	"github.com/gin-gonic/gin"
)

func GinInit(){
	r := gin.Default()
	r = router(r)
	_ = r.Run("0.0.0.0:"+config.GetString("server_port"))
}