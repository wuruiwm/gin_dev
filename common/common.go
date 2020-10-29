package common

import (
	"crypto/md5"
	"encoding/hex"
	"gin_dev/config"
	"math/rand"
	"time"
)

//获取带协议头的域名
func GetDomainName()string{
	return config.GetString("domain_name")
}

//是否开启debug
func IsDebug()bool{
	return config.GetBool("debug")
}

//获取app_key
func GetAppKey()string{
	return config.GetString("app_key")
}

//md5计算
func MD5(str string)string{
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//生成随机字符串
func GetRandString(num int)string{
	rand.Seed(time.Now().UnixNano())
	var rangString string
	strpol := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i:=0;i<num;i++ {
		tmpNum := rand.Intn(len(strpol)-1)
		rangString += strpol[tmpNum:tmpNum+1]
	}
	return rangString
}

//获取时间戳(秒)
func GetUnixTime()int{
	return int(time.Now().Unix())
}