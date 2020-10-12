package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
func (r Response)Success(c *gin.Context,msg string,data interface{}){
	r.Return(c,SuccessCode,msg,data)
}

//失败返回
func (r Response)Error(c *gin.Context,msg string){
	r.Return(c,ErrorCode,msg,nil)
}

//通用返回
func (r Response)Return(c *gin.Context,code int,msg string,data interface{}){
	r.Code = code
	r.Msg = msg
	r.Data = data
	r.JSON(c)
}

//通用json返回
func (r Response)JSON(c *gin.Context){
	c.JSON(http.StatusOK,r)
}