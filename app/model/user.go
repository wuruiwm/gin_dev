package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin_dev/app/response"
	"gin_dev/common"
)

type User struct {
	ID         uint   `gorm:"column:id;not null;primary_key;AUTO_INCREMENT;type:int(11)" json:"id"`
	Username   string `gorm:"column:username;not null;default:'';index;type:char(50)" json:"username"`
	Password   string `gorm:"column:password;not null;default:'';comment:'hash后的密码';type:char(255)" json:"password"`
	Salt       string `gorm:"column:salt;not null;default:'';comment:'hash密码的时候加入的salt';type:char(50)" json:"salt"`
	CreateTime int    `gorm:"column:create_time;not null;default:0;comment:'创建时间';type:int(11)" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time;not null;default:0;comment:'修改时间';type:int(11)" json:"update_time"`
	LastLoginTime int `gorm:"column:last_login_time;not null;default:0;comment:'最后登陆时间';type:int(11)" json:"last_login_time"`
}

func (*User)TableName()string{
	return `user`
}

type LoginResult struct {
	Token string `json:"token"`
	Expire int   `json:"expire"`
}

func AdminLogin(username string,password string)(*LoginResult,error){
	var (
		user *User
		err error
		loginResult LoginResult
	)
	if user,err = CheckUserPassword(username,password);err != nil{
		return &loginResult,err
	}
	return LoginSuccess(user)
}

func CheckUserPassword(username string,password string)(*User,error){
	var (
		user User
		err error
	)
	if err = db.Where("username",username).Select("id,username,password,salt").Take(&user).Error;err != nil{
		return &user,errors.New("用户不存在")
	}
	if GetHashPassword(password,user.Salt) != user.Password{
		return &user,errors.New("密码不正确")
	}
	return &user,nil
}

func GetHashPassword(password string,salt string)string{
	appKey := common.GetAppKey()
	return common.MD5(common.MD5(password+appKey)+salt)
}

func LoginSuccess(user *User)(*LoginResult,error){
	var (
		loginResult LoginResult
		err error
	)
	//更新登陆时间
	user.LastLoginTime = common.GetUnixTime()
	if err = db.Select("last_login_time").Updates(user).Error;err != nil{
		return &loginResult,errors.New("修改最后登陆时间失败")
	}
	return SetToken(user)
}

func SetToken(user *User)(*LoginResult,error){
	var (
		loginResult LoginResult
		cacheKey string
		buf []byte
		err error
	)
	//设置token
	loginResult.Token = GetHashToken(user)
	loginResult.Expire = 60*60*24*7
	cacheKey = GetAdminTokenPrefix() + loginResult.Token
	buf,err = json.Marshal(user)
	if err != nil{
		return &loginResult,errors.New("json序列化失败")
	}
	err = common.SetCache(cacheKey,string(buf),loginResult.Expire)
	return &loginResult,err
}

func GetAdminTokenPrefix()string{
	return "admin:token:"
}

func GetTokenUser(token string)(*User,int,error){
	var (
		cacheKey string
		userJson string
		user User
		err error
	)
	if token == ""{
		return &user,response.NotLoginCode,errors.New("未登录")
	}
	cacheKey = GetAdminTokenPrefix() + token
	userJson = common.GetCache(cacheKey)
	if userJson == ""{
		return &user,response.LoginExpiredCode,errors.New("登陆已过期")
	}
	err = json.Unmarshal([]byte(userJson),&user)
	return &user,response.ErrorCode,err
}

func GetUser(userId uint)(*User,error){
	var (
		user User
		err error
	)
	if err = db.Where("id",userId).Take(&user).Error;err != nil{
		err = errors.New("用户不存在")
	}
	return &user,err
}

func GetHashToken(user *User)string{
	return common.MD5(fmt.Sprintf("%s%s%s%d",user.Username,user.Password,user.Salt,user.LastLoginTime))
}