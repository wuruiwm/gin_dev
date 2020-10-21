package model

import (
	"fmt"
	"gin_dev/common"
	"gin_dev/config"
	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var Db *gorm.DB
var Redis *redis.Pool

//初始化操作
func init(){
	mysqlInit()
	redisInit()
}

//初始化mysql连接池
func mysqlInit(){
	var gormConfig *gorm.Config
	if common.IsDebug(){
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),//打印所有执行sql
		}
	}else{
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),//只打印错误sql
		}
	}
	conn, err := gorm.Open(mysql.Open(getMysqlConnString()),gormConfig)
	if err != nil{
		fmt.Println("mysql连接错误:",err)
		os.Exit(1)
	}
	Db = conn
}

//获取mysql连接字符串
func getMysqlConnString()string{
	username := config.GetString("db.mysql.username")
	password := config.GetString("db.mysql.password")
	host := config.GetString("db.mysql.host")
	port := config.GetString("db.mysql.port")
	dbname := config.GetString("db.mysql.dbname")
	charset := config.GetString("db.mysql.charset")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",username,password,host,port,dbname,charset)
}

//初始化redis连接池
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

//获取redis连接字符串
func getRedisConnString()string{
	host := config.GetString("db.redis.host")
	port := config.GetString("db.redis.port")
	return fmt.Sprintf("%s:%s",host,port)
}