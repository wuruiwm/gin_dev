package model

import (
	"fmt"
	"gin_dev/config"
	"github.com/garyburd/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB
var Redis *redis.Pool

func init(){
	mysqlInit()
	redisInit()
}

func mysqlInit(){
	conn, err := gorm.Open(mysql.Open(getMysqlConnString()), &gorm.Config{})
	if err != nil{
		fmt.Println("mysql连接错误:",err)
		os.Exit(1)
	}
	Db = conn
}

func getMysqlConnString()string{
	return config.GetString("db.mysql.user")+":"+config.GetString("db.mysql.passwd")+"@tcp("+config.GetString("db.mysql.host")+":"+config.GetString("db.mysql.port")+")/"+config.GetString("db.mysql.name")+"?charset="+config.GetString("db.mysql.charset")+"&parseTime=True&loc=Local"
}

func redisInit(){
	Redis = &redis.Pool{
		MaxIdle:   4,
		MaxActive: 16,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", getRedisConnString())
			if err != nil {
				fmt.Println("redis连接错误:",err)
				os.Exit(1)
			}
			return c, err
		},
	}
}

func getRedisConnString()string{
	return config.GetString("db.redis.host") + ":" + config.GetString("db.redis.port")
}