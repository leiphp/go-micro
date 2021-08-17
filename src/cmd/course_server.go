package main

import (
	"github.com/micro/go-micro/v2"
	"go-micro/src/Boot"
	"go-micro/src/Course"
	"go-micro/src/service"
	"log"
)


func main(){
	//加载各种配置，初始化等
	Boot.BootInit()
	cService := micro.NewService(
		micro.Name("go.micro.api.course"))
	cService.Init()
	err := Course.RegisterCourseServiceHandler(cService.Server(), service.NewCourseServiceImpl())
	if err != nil {
		log.Fatal(err)
	}
	if err = cService.Run(); err != nil {
		log.Println(err)
	}
}
