package lesson

import (
	lessonEntity "go-management-auth-school/entity/lesson"
)

var (
	InsertLesson = `INSERT INTO ` + lessonEntity.Table + ` (id, name, type) VALUES (?,?,?)`
)
