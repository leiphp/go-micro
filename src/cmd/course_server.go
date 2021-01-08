package main

import (
	"github.com/micro/go-micro/v2"
	"go-micro/src/Course"
	"log"
)


func main(){
	service := micro.NewService(
		micro.Name("api.100txy.com.course"))
	service.Init()
	err := Course.RegisterCourseServiceHandler(service.Server(), Course.NewCourseServiceImpl())
	if err != nil {
		log.Fatal(err)
	}
	if err = service.Run(); err != nil {
		log.Println(err)
	}
}
