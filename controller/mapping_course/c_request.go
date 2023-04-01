package mapping_course

import "go-management-auth-school/controller"


type MappingCourseParams struct {
	CourseID int `json:"course_id"`
	TeacherID int `json:"teacher_id"`
	LessonID int `json:"lesson_id"`
	controller.DefaultParameter
}