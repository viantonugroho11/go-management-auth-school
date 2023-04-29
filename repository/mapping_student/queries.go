package mapping_student

import (
	mapStudentEntity "go-management-auth-school/entity/mapping_student"
)

var (
	InsertMapStudent = `INSERT INTO ` + mapStudentEntity.Table + ` (id,identity_id, class_id, type) VALUES (?,?, ?, ?)`
)
