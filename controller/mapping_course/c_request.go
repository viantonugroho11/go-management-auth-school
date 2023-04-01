package mapping_course

import "go-management-auth-school/controller"


type MappingCourseParams struct {
	ClassID  string `json:"class_id"`
	TeacherID string `json:"teacher_id"`
	LessonID string `json:"lesson_id"`
	controller.DefaultParameter
}