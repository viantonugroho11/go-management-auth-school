package parent

import (
	parentEntity "go-management-auth-school/entity/parent"
)

var (
	InsertParent = `INSERT INTO ` + parentEntity.TableName + ` (
		id,
		first_name,
		last_name,
		type,
		nik,
		gender,
		phone,
		work_id,
		work_name,
		income,
		student_id,
		image
	) VALUES (
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?
	)`
)
