package middleware

import (
	"gin_dev/app/model"
	"gin_dev/app/response"
	"github.com/gin-gonic/gin"
)

func AdminAuth()gin.HandlerFunc{
	return func(c *gin.Context){
		var responseCode int
		defer func() {
			if err := recover();err != nil{
				if msg,ok := err.(string);ok{
					response.Return(c,responseCode,msg,nil)
				}else{
					response.Error(c,"用户鉴权失败")
				}
				c.Abort()
			}
		}()
		//获取token
		token := c.GetHeader("token")
		userCache,responseCode,err := model.GetTokenUser(token)
		if err != nil{
			panic(err.Error())
		}
		//获取用户表数据
		user,err := model.GetUser(userCache.ID)
		if err != nil{
			responseCode = response.NotUserCode
			panic(err.Error())
		}
		//校验密码是否被修改
		if user.Password != userCache.Password{
			responseCode = response.PasswordChangeCode
			panic("密码被修改,请重新登陆")
		}
		//单点登录校验
		if model.GetHashToken(user) != model.GetHashToken(userCache){
			responseCode = response.OtherLoginCode
			panic("您的账号在其他地方登陆")
		}
		//将id 用户名存起来 以供上下文使用
		c.Set("user_id",int(user.ID))
		c.Set("username",user.Username)

		c.Next()
	}
}