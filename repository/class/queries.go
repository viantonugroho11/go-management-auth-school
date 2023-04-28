package class



import (
	classEntity "go-management-auth-school/entity/class"
)

var (
	InsertClass = `INSERT INTO ` + classEntity.Table + ` (id, name, major_id, level) VALUES (?,?,?,?)`
)

