package mapping_course

import (
	classEntity "go-management-auth-school/entity/class"
	lessonEntity "go-management-auth-school/entity/lesson"
	teacherEntity "go-management-auth-school/entity/teacher"
)

type MappingCourse struct {
	ID        int                   `json:"id"`
	ClassID   string                `json:"class_id"`
	TeacherID string                `json:"teacher_id"`
	LessonID  string                `json:"lesson_id"`
	Class     classEntity.Class     `json:"class"`
	Teacher   teacherEntity.Teacher `json:"teacher"`
	Lesson    lessonEntity.Lesson   `json:"lesson"`
}
