package main

import (
	_ "gin_dev/config"
	_ "gin_dev/model"
	"gin_dev/router"
)

func main(){
	router.HttpInit()
}