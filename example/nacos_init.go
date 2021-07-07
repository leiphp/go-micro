package main

import (
	"fmt"
	"go-micro/src/Config"
	"time"
)

//封装初始化去nacos拿配置
func main(){
	Config.InitConfig()
	fmt.Println(Config.JConfig.Config)
	fmt.Println(Config.JConfig.Data.Mysql)

	for{
		time.Sleep(time.Second*1)
		fmt.Println(Config.JConfig.Data.Mysql)
	}
}