package main

import (
	"github.com/micro/go-micro/v2"
	"go-micro/src/Course"
	"go-micro/src/service"
	"log"
)


func main(){
	cService := micro.NewService(
		micro.Name("api.100txy.com.course"))
	cService.Init()
	err := Course.RegisterCourseServiceHandler(cService.Server(), service.NewCourseServiceImpl())
	if err != nil {
		log.Fatal(err)
	}
	if err = cService.Run(); err != nil {
		log.Println(err)
	}
}
