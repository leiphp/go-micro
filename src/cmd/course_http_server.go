package main

import (
	"context"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/web"
	"go-micro/src/Course"
	"github.com/gin-gonic/gin"
	"log"
)


func main(){
	//go-micro集成gin开发http api
	c := grpc.NewClient()
	r := gin.New()
	r.Handle("GET","/test", func(ctx *gin.Context){
		c := Course.NewCourseService("api.100txy.com.course",c)
		course_rsp,_ := c.ListForTop(context.Background(),&Course.ListRequest{Size:10})
		ctx.JSON(200, gin.H{"Result":course_rsp.Result})
	})

	service := web.NewService(
		web.Name("api.100txy.com.http.course"),
		web.Handler(r),
		)
	//service.HandleFunc("/test",func(writer http.ResponseWriter, request *http.Request){
	//	c := Course.NewCourseService("api.100txy.com.course",c)
	//	course_rsp,_ := c.ListForTop(context.Background(),&Course.ListRequest{Size:10})
	//	log.Println(course_rsp.Result)
	//	writer.Write([]byte("http api test"))
	//})
	service.Init()

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
