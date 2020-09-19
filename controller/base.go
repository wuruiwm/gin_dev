package controller

import "github.com/gin-gonic/gin"

func success(c *gin.Context,msg string,data interface{},othor map[string]interface{}){
	response := gin.H{
		"code":0,
		"msg":msg,
	}
	if data != nil{
		response["data"] = data
	}
	if othor != nil{
		for k,v := range othor{
			response[k] = v
		}
	}
	json(c,response)
}

func error(c *gin.Context,msg string){
	response := gin.H{
		"code":1,
		"msg":msg,
	}
	json(c,response)
}

func json(c *gin.Context,response interface{}){
	c.JSON(200,response)
}