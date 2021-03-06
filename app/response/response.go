package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode = 0 //通用成功 状态码
	ErrorCode = 1 //通用错误 状态码

	NotLoginCode = 40001 //未登录
	LoginExpiredCode = 40002 //登陆已过期
	PasswordChangeCode = 40003 //密码被修改
	NotUserCode = 40004 //用户不存在
	OtherLoginCode = 40005 //在其他地方登陆
)

//http response struct
type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

//通用成功返回
func Success(c *gin.Context,msg string,data interface{}){
	Return(c,SuccessCode,msg,data)
}

//通用失败返回
func Error(c *gin.Context,msg string){
	Return(c,ErrorCode,msg,nil)
}

//通用返回
func Return(c *gin.Context,code int,msg string,data interface{}){
	r := &Response{
		Code:code,
		Msg:msg,
		Data:data,
	}
	JSON(c,r)
}

//通用json返回
func JSON(c *gin.Context,r *Response){
	c.JSON(http.StatusOK,r)
}

//控制器内捕获错误 并以通用失败返回
func RecoverError(c *gin.Context,defaultMsg string){
	if err := recover();err != nil{
		if msg,ok := err.(string);ok{
			Error(c,msg)
		}else{
			Error(c,defaultMsg)
		}
	}
}