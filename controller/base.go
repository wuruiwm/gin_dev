package controller

import (
	"github.com/gin-gonic/gin"
)

const SuccessCode int = 0 //成功 状态码
const ErrorCode int = 1 //通用错误 状态码

//http response struct
type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

//成功返回
func (r *Response)Success(c *gin.Context,msg string,data interface{}){
	response := &Response{
		Code:SuccessCode,
		Msg:msg,
		Data:data,
	}
	response.JSON(c)
}

//失败返回
func (r *Response)Error(c *gin.Context,msg string){
	response := &Response{
		Code:ErrorCode,
		Msg:msg,
	}
	response.JSON(c)
}

//通用json
func (r *Response)JSON(c *gin.Context){
	c.JSON(200,r)
}