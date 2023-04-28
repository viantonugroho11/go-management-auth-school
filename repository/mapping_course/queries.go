package mapping_course

import (
	mapCourseEntity "go-management-auth-school/entity/mapping_course"
)

var (
	InsertMapCourse = `INSERT INTO ` + mapCourseEntity.Table + ` (id, class_id, teacher_id, lesson_id) VALUES (?,?,?,?)`
)
