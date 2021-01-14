package service

import (
	"context"
	"go-micro/src/Boot"
	."go-micro/src/Course"
	"go-micro/src/Vars"
)

func NewCourseModel (id int32, name string) *CourseModel {
	return &CourseModel{CourseId:id,CourseName:name}
}

type CourseServiceImpl struct {

}

func (this *CourseServiceImpl) ListForTop(ctx context.Context, req *ListRequest, rsp *ListResponse)  error {
	//ret := make([]*CourseModel,0)
	//ret = append(ret,NewCourseModel(101,"java课程"),NewCourseModel(102,"php课程"))
	//rsp.Result = ret
	//return nil

	course := make([]*CourseModel,0)
	err := Boot.GetDB().Table(Vars.Table_CourseMain).Order("course_id desc").Find(&course).Error
	if err != nil {
		return err
	}
	rsp.Result = course
	return nil
}

func NewCourseServiceImpl() *CourseServiceImpl {
	return &CourseServiceImpl{}
}
