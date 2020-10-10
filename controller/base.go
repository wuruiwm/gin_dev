package controller

import "github.com/gin-gonic/gin"

const SuccessCode int = 0 //成功 状态码
const ErrorCode int = 1 //通用错误 状态码

type Response struct {
	code int
	msg string
	data interface{}
}

func (r *Response)Success(c *gin.Context,msg string,data interface{}){
	r.code = SuccessCode
	r.msg = msg
	r.data = data
	r.JSON(c)
}

func (r *Response)Error(c *gin.Context,msg string){
	r.code = ErrorCode
	r.msg = msg
	r.JSON(c)
}

func (r *Response)JSON(c *gin.Context){
	c.JSON(200,r)
}