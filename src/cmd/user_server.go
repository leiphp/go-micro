package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	client2 "github.com/micro/go-micro/v2/client"
	"go-micro/src/Course"
	"go-micro/src/Users"
	"log"
)

type UserService struct {
	client client2.Client
}

func(this *UserService) Test(ctx context.Context, req *Users.UserRequest, rsp *Users.UserResponse) error {
	rsp.Ret = "users" + req.Id
	//服务之间调用
	c := Course.NewCourseService("go.micro.api.course",this.client)
	course_rsp,_ := c.ListForTop(ctx,&Course.ListRequest{Size:10})
	log.Println(course_rsp.Result)
	return nil
}

func NewUserService(c client2.Client) *UserService {
	return &UserService{client:c}
}

func main(){
	service := micro.NewService(
		micro.Name("go.micro.api.user"))
	service.Init()

	err := Users.RegisterUserServiceHandler(service.Server(), NewUserService(service.Client()))
	if err != nil {
		log.Fatal(err)
	}
	if err = service.Run(); err != nil {
		log.Println(err)
	}
}
