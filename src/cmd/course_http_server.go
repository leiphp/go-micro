package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"go-micro/framework/gin_"
	"go-micro/src/Boot"
	_ "go-micro/src/lib"
	"strings"
)


func main(){
	//go-micro集成gin开发http api
	//c := grpc.NewClient()
	Boot.BootInit()	//加载各种配置，初始化等
	r := gin.New()
	//r.Handle("GET","/test", func(ctx *gin.Context){
	//	c := Course.NewCourseService("api.100txy.com.course",c)
	//	course_rsp,_ := c.ListForTop(context.Background(),&Course.ListRequest{Size:10})
	//	ctx.JSON(200, gin.H{"Result":course_rsp.Result})
	//})

	gin_.BootStrap(r)

	service := web.NewService(
		web.Name(strings.Join([]string{Boot.JConfig.Service.Namespace,Boot.JConfig.Service.Name},".")),
		web.Handler(r),
		)
	service.Init()
	go func() {
		if err := service.Run(); err != nil {
			Boot.BootErrChan<-err
		}
	}()
	getErr := <-Boot.BootErrChan
	logger.Info(getErr)

}
