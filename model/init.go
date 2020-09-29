package model

import (
	"fmt"
	"gin_dev/config"
	"github.com/gomodule/redigo/redis"
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
	username := config.GetString("db.mysql.username")
	password := config.GetString("db.mysql.password")
	host := config.GetString("db.mysql.host")
	port := config.GetString("db.mysql.port")
	dbname := config.GetString("db.mysql.dbname")
	charset := config.GetString("db.mysql.charset")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",username,password,host,port,dbname,charset)
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
	host := config.GetString("db.redis.host")
	port := config.GetString("db.redis.port")
	return fmt.Sprintf("%s:%s",host,port)
}