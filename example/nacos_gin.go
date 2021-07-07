package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-micro/src/Config"
	"log"
	"time"
)

//网址服务读取配置，延迟处理
func main(){
	Config.InitConfig()
	fmt.Println(Config.JConfig.Data.Mysql)

	err := WaitForReady()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Handle("GET","/", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":Config.JConfig.Data.Mysql})
	})
	r.Run(":8111")
}


func WaitForReady() error {
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*5)//允许秒数
	defer cancel()
	for{
		select {
		case <-ctx.Done():
			return fmt.Errorf("init config error")
		default:
			if Config.JConfig.Data.Mysql != nil {
				return nil
			}
		}
	}
}