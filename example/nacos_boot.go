package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-micro/src/Config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//网址服务读取配置，延迟处理
func main(){
	Config.InitConfig()
	fmt.Println(Config.JConfig.Data.Mysql)
	errChan := make(chan error)
	go func() {
		err := WaitForReady4()
		if err != nil {
			errChan<-err
		}
	}()



	r := gin.New()
	r.Use(CheckForReady4())
	r.Handle("GET","/", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":Config.JConfig.Data.Mysql})
	})

	go func() {
		err := r.Run(":8111")
		if err != nil {
			errChan<-err
		}
	}()

	go func() {
		sig_c := make(chan os.Signal)
		signal.Notify(sig_c,syscall.SIGINT,syscall.SIGTERM)
		errChan<-fmt.Errorf("%s",<-sig_c)
	}()

	getErr := <-errChan
	log.Fatal(getErr)

}


func CheckForReady4() gin.HandlerFunc {
	return func(i *gin.Context) {
		if Config.JConfig.Data.Mysql == nil {
			i.AbortWithStatusJSON(200,gin.H{"result":"server is loading"})
		}else{
			i.Next()
		}
	}
}

func WaitForReady4() error {
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*8)//允许秒数
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