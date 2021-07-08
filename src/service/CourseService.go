package service

import (
	"context"
	. "go-micro/src/Course"
	"go-micro/src/mapper"
)
//grpc和http公共服务文件
func NewCourseModel (id int32, name string) *CourseModel {
	return &CourseModel{CourseId:id,CourseName:name}
}

type CourseServiceImpl struct {

}

//获取置顶课程列表
func (this *CourseServiceImpl) ListForTop(ctx context.Context, req *ListRequest, rsp *ListResponse)  error {
	course := make([]*CourseModel,0)
	err := mapper.GetCourseListBySql(1).Find(&course).Error
	if err != nil {
		return err
	}
	rsp.Result = course
	return nil
}

//获取课程详情
func (this *CourseServiceImpl) GetDetail(ctx context.Context, req *DetailRequest, rsp *DetailResponse)  error {

	if err := mapper.GetCourseDetail(int(req.CourseId)).Find(rsp.Result).Error; err != nil{
		return err
	}
	return nil
}

func NewCourseServiceImpl() *CourseServiceImpl {
	return &CourseServiceImpl{}
}
