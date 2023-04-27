package mapping_course

import (
	"fmt"
	"go-management-auth-school/controller"

	"github.com/go-playground/validator"
	mpCourseEntity "go-management-auth-school/entity/mapping_course"
)

type MappingCourseParams struct {
	ClassID   string `json:"class_id"`
	TeacherID string `json:"teacher_id"`
	LessonID  string `json:"lesson_id"`
	controller.DefaultParameter
}

type MappingCourseRequest struct {
	ClassID   string `json:"class_id"`
	TeacherID string `json:"teacher_id"`
	LessonID  string `json:"lesson_id"`
}

func (m *MappingCourseRequest) Validate() error {
	err := validator.New().Struct(m)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}
	return nil
}

func (m *MappingCourseRequest) ToEntity() (res *mpCourseEntity.MappingCourse) {
	res = &mpCourseEntity.MappingCourse{
		ClassID:   m.ClassID,
		TeacherID: m.TeacherID,
		LessonID:  m.LessonID,
	}
	return
}
