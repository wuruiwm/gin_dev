package model

import (
	"errors"
	"fmt"
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

func (*User) TableName() string {
	return `user`
}

func AdminLogin(username string,password string)error{
	var (
		user *User
		err error
	)
	if user,err = CheckUserPassword(username,password);err != nil{
		return err
	}
	fmt.Println(user)
	LoginSuccess(user)
	return nil
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

func LoginSuccess(user *User){
	//更新登陆时间
	user.LastLoginTime = common.GetUnixTime()
	db.Where("id",user.ID).Updates(user)
}