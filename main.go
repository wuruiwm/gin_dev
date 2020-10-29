package main

import (
	_ "gin_dev/app/model"
	_ "gin_dev/common"
	_ "gin_dev/config"
	"gin_dev/router"
)

func main(){
	router.HttpInit()
}