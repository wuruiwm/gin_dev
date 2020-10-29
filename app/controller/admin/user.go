package admin

import (
	"gin_dev/app/model"
	"gin_dev/app/response"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	defer response.RecoverError(c,"登陆出错")
	//获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	//参数校验
	if username == ""{
		panic("用户名不能为空")
	}
	if password == ""{
		panic("密码不能为空")
	}

	//校验用户名密码是否正确 并登陆
	loginResult,err := model.AdminLogin(username,password)
	if err != nil{
		panic(err.Error())
	}
	response.Success(c,"登陆成功",loginResult)
}