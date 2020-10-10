package controller

import "github.com/gin-gonic/gin"

const SuccessCode int = 0 //成功 状态码
const ErrorCode int = 1 //通用错误 状态码

func Success(c *gin.Context,msg string,data interface{},other map[string]interface{}){
	response := gin.H{
		"code":SuccessCode,
		"msg":msg,
	}
	if data != nil{
		response["data"] = data
	}
	if other != nil{
		for k,v := range other {
			response[k] = v
		}
	}
	JSON(c,response)
}

func Error(c *gin.Context,msg string){
	response := gin.H{
		"code": ErrorCode,
		"msg":  msg,
	}
	JSON(c,response)
}

func JSON(c *gin.Context,response interface{}){
	c.JSON(200,response)
}