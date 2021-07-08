package service

import (
	"context"
	"fmt"
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
	//只取课程详细
	fmt.Println("fetch_type",req.FetchType)
	if req.FetchType==0 || req.FetchType==1|| req.FetchType==3{
		if err := mapper.GetCourseDetail(int(req.CourseId)).Find(rsp.Course).Error; err != nil{
			return err
		}
	}
	//只取计数表详情
	if req.FetchType==2||req.FetchType==3{
		if err := mapper.GetCourseCounts(int(req.CourseId)).Find(&rsp.Counts).Error; err != nil{
			return err
		}
	}

	return nil
}

func NewCourseServiceImpl() *CourseServiceImpl {
	return &CourseServiceImpl{}
}
