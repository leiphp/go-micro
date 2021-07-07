package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	"go-micro/src/Boot"
	"go-micro/src/Config"
	"os"
	"os/signal"
	"syscall"
)

//网址服务读取配置，延迟处理
func main(){
	Boot.BootInit()
	r := gin.New()
	r.Use(CheckForReady3())
	r.Handle("GET","/", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":Config.JConfig.Data.Mysql})
	})

	go func() {
		err := r.Run(":8111")
		if err != nil {
			Boot.BootErrChan<-err
		}
	}()

	go func() {
		sig_c := make(chan os.Signal)
		signal.Notify(sig_c,syscall.SIGINT,syscall.SIGTERM)
		Boot.BootErrChan<-fmt.Errorf("%s",<-sig_c)
	}()

	getErr := <-Boot.BootErrChan
	logger.Info(getErr)

}


func CheckForReady3() gin.HandlerFunc {
	return func(i *gin.Context) {
		if Config.JConfig.Data.Mysql == nil {
			i.AbortWithStatusJSON(200,gin.H{"result":"server is loading"})
		}else{
			i.Next()
		}
	}
}
