package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

//初始化配置文件
func init(){
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败:",err)
		os.Exit(1)
	}
}

func GetString(key string)string{
	return viper.GetString(key)
}

func GetBool(key string)bool{
	return viper.GetBool(key)
}

func GetInt(key string)int{
	return viper.GetInt(key)
}