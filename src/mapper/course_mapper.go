package mapper

import (
	"github.com/jinzhu/gorm"
	"go-micro/src/Boot"
	"go-micro/src/vars"
)

//gorm获取列表
func GetCourseList() *gorm.DB{
	return Boot.GetDB().Table(vars.Table_CourseMain).Order("course_id desc").Limit(3)
}

//原生sql获取列表
const course_list="select * from course_main order by course_id desc limit ?"
func GetCourseListBySql(args ...interface{}) *gorm.DB{
	return Boot.GetDB().Raw(course_list,args...)
}

//获取课程详情
func GetCourseDetail(course_id int) *gorm.DB{
	return Boot.GetDB().Table(vars.Table_CourseMain).Where("course_id = ?",course_id)
}
