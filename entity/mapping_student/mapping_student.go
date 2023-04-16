package mappingstudent

import (
	studentEntity "go-management-auth-school/entity/student"
	classEntity "go-management-auth-school/entity/class"
	teacherEntity "go-management-auth-school/entity/teacher"
	mpCourseEntity "go-management-auth-school/entity/mapping_course"
)

type MappingStudent struct {
	ID int `json:"id"`
	Indentity string `json:"indentity"`
	ClassID string `json:"class_id"`
	Type string `json:"type"`
	Student studentEntity.Student `json:"student"`
	Class classEntity.Class `json:"class"`
	Teacher teacherEntity.Teacher `json:"teacher"`
	MpCourse []mpCourseEntity.MappingCourse `json:"mp_course"`
}

var (

	Table = "mapping_student"

	Column = []string{"def.id", "def.indentity_id", "def.class_id", "def.type"}

)