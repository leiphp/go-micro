package main

import (
	"github.com/micro/go-micro/v2/web"
	"go-micro/framework/gin_"
	_ "go-micro/src/lib"
	"github.com/gin-gonic/gin"
	"log"
)


func main(){
	//go-micro集成gin开发http api
	//c := grpc.NewClient()
	r := gin.New()
	//r.Handle("GET","/test", func(ctx *gin.Context){
	//	c := Course.NewCourseService("api.100txy.com.course",c)
	//	course_rsp,_ := c.ListForTop(context.Background(),&Course.ListRequest{Size:10})
	//	ctx.JSON(200, gin.H{"Result":course_rsp.Result})
	//})

	gin_.BootStrap(r)

	service := web.NewService(
		web.Name("api.100txy.com.http.course"),
		web.Handler(r),
		)
	service.Init()

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
