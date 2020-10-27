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
	"time"
)

var db *gorm.DB
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
	db = conn
	setMysqlSetting()
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
		MaxActive: 8,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", getRedisConnString())
			if err != nil {
				fmt.Println("redis连接错误:",err)
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

//设置mysql连接池参数
func setMysqlSetting(){
	sqlDB,err := db.DB()
	if err != nil{
		fmt.Println("mysql设置连接池参数错误:",err)
		os.Exit(1)
	}
	//用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(4)
	//设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(8)
	//设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetRedisConn()redis.Conn{
	return Redis.Get()
}