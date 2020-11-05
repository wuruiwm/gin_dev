package common

import (
	"fmt"
	"gin_dev/config"
	"github.com/gomodule/redigo/redis"
)

var Redis *redis.Pool

//获取redis连接字符串
func getRedisConnString()string{
	host := config.GetString("db.redis.host")
	port := config.GetString("db.redis.port")
	return fmt.Sprintf("%s:%s",host,port)
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

//获取一个redis连接
func GetRedisConn()redis.Conn{
	return Redis.Get()
}

//设置缓存
func SetCache(key string,data string,expire int)error{
	var err error
	conn := GetRedisConn()
	defer conn.Close()
	if expire == 0{
		_,err = conn.Do("set",key,data)
	}else{
		_,err = conn.Do("setex",key,expire,data)
	}
	return err
}

//获取缓存
func GetCache(key string)string{
	conn := GetRedisConn()
	defer conn.Close()
	r,err := redis.String(conn.Do("get", key))
	if err != nil{
		return ""
	}
	return r
}